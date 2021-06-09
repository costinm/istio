// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc

import (
	"context"
	"io"
	"log"
	"strings"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"istio.io/istio/pilot/pkg/features"
	istiokeepalive "istio.io/istio/pkg/keepalive"
)

type SendHandler func() error

var timeout = features.XdsPushSendTimeout

// Send with timeout if specified. If timeout is zero, sends without timeout.
func Send(ctx context.Context, send SendHandler) error {
	if timeout.Nanoseconds() > 0 {
		errChan := make(chan error, 1)
		timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		go func() {
			err := send()
			errChan <- err
			close(errChan)
		}()
		select {
		case <-timeoutCtx.Done():
			return status.Errorf(codes.DeadlineExceeded, "timeout sending")
		case err := <-errChan:
			return err
		}
	}
	err := send()
	return err
}
func init() {
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Println("Failed to register ocgrpc server views: %v", err)
	}
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Println("Failed to register ocgrpc server views: %v", err)
	}
}
func ServerOptions(options *istiokeepalive.Options, interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	maxStreams := features.MaxConcurrentStreams
	maxRecvMsgSize := features.MaxRecvMsgSize

	grpcOptions := []grpc.ServerOption{
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(interceptors...)),
		grpc.MaxConcurrentStreams(uint32(maxStreams)),
		grpc.MaxRecvMsgSize(maxRecvMsgSize),
		// Ensure we allow clients sufficient ability to send keep alives. If this is higher than client
		// keep alive setting, it will prematurely get a GOAWAY sent.
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: options.Time / 2,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:                  options.Time,
			Timeout:               options.Timeout,
			MaxConnectionAge:      options.MaxServerConnectionAge,
			MaxConnectionAgeGrace: options.MaxServerConnectionAgeGrace,
		}),
	}

	return grpcOptions
}

// IsExpectedGRPCError checks a gRPC error code and determines whether it is an expected error when
// things are operating normally. This is basically capturing when the client disconnects.
func IsExpectedGRPCError(err error) bool {
	if err == io.EOF {
		return true
	}

	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.Canceled || s.Code() == codes.DeadlineExceeded {
			return true
		}
		if s.Code() == codes.Unavailable && (s.Message() == "client disconnected" || s.Message() == "transport is closing") {
			return true
		}
	}
	// If this is not a gRPCStatus we should just error message.
	if strings.Contains(err.Error(), "stream terminated by RST_STREAM with error code: NO_ERROR") {
		return true
	}

	return false
}
