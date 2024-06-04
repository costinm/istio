local g = import 'g.libsonnet';

local row = g.panel.row;

local grid = import 'lib-grid.libsonnet';
local panels = import './panels.libsonnet';
local variables = import './variables.libsonnet';
local queries = import './queries.libsonnet';

g.dashboard.new('Istio Ztunnel Dashboard')
+ g.dashboard.graphTooltip.withSharedCrosshair()
+ g.dashboard.withVariables([
  variables.datasource,
])
+ g.dashboard.withPanels(
  grid.makeGrid([
    row.new('Process')
    + row.withPanels([
      panels.timeSeries.base('Ztunnel Versions', queries.istioBuild, 'Version number of each running instance'),
      panels.timeSeries.bytes('Memory Usage', queries.memUsage, 'Memory usage of each running instance'),
      panels.timeSeries.base('CPU Usage', queries.cpuUsage, 'CPU usage of each running instance'),
    ]),
    row.new('Network')
    + row.withPanels([
      panels.timeSeries.connections('Connections', queries.connections, 'Connections opened and closed per instance'),
      panels.timeSeries.bytesRate('Bytes Transmitted', queries.bytes, 'Bytes sent and recieved per instance'),
      panels.timeSeries.dns('DNS Request', queries.dns, 'DNS queries recieved per instance'),
    ]),
    row.new('Operations')
    + row.withPanels([
      panels.timeSeries.base(
        'XDS', queries.xdsConnections, |||
          Count of XDS connection terminations.
          This will typically spike every 30min for each instance.
        |||
      ),
      panels.timeSeries.base('Workload Manager', queries.workloadManager, |||
        Count of active and pending proxies managed by each instance.
        Pending is expected to converge to zero.
      |||
      ),
    ]),
  ], panelHeight=8)
)
+ g.dashboard.withUid(std.md5('ztunnel.json'))
