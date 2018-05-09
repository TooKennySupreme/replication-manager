// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017 Signal 18 SARL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.

package cluster

func (cluster *Cluster) SwitchServerMaintenance(serverid uint) {
	server := cluster.GetServerFromId(serverid)
	server.SwitchMaintenance()
	cluster.SetProxyServerMaintenance(server.ServerID)
}
func (cluster *Cluster) SwitchProvNetCNI() {
	cluster.Conf.ProvNetCNI = !cluster.Conf.ProvNetCNI
}
func (cluster *Cluster) SwitchProvDockerDaemonPrivate() {
	cluster.Conf.ProvDockerDaemonPrivate = !cluster.Conf.ProvDockerDaemonPrivate
}

func (cluster *Cluster) SwitchInteractive() {
	if cluster.Conf.Interactive == true {
		cluster.Conf.Interactive = false
		cluster.LogPrintf(LvlInfo, "Failover monitor switched to automatic mode")
	} else {
		cluster.Conf.Interactive = true
		cluster.LogPrintf(LvlInfo, "Failover monitor switched to manual mode")
	}
}

func (cluster *Cluster) SwitchReadOnly() {
	cluster.Conf.ReadOnly = !cluster.Conf.ReadOnly
}

func (cluster *Cluster) SwitchRplChecks() {
	cluster.Conf.RplChecks = !cluster.Conf.RplChecks
}

func (cluster *Cluster) SwitchCleanAll() {
	cluster.CleanAll = !cluster.CleanAll
}

func (cluster *Cluster) SwitchFailSync() {
	cluster.Conf.FailSync = !cluster.Conf.FailSync
}

func (cluster *Cluster) SwitchSwitchoverSync() {
	cluster.Conf.SwitchSync = !cluster.Conf.SwitchSync
}

func (cluster *Cluster) SwitchVerbosity() {

	if cluster.GetLogLevel() > 0 {
		cluster.SetLogLevel(0)
	} else {
		cluster.SetLogLevel(4)
	}
}

func (cluster *Cluster) SwitchRejoin() {
	cluster.Conf.Autorejoin = !cluster.Conf.Autorejoin
}

func (cluster *Cluster) SwitchRejoinDump() {
	cluster.Conf.AutorejoinMysqldump = !cluster.Conf.AutorejoinMysqldump
}

func (cluster *Cluster) SwitchRejoinBackupBinlog() {
	cluster.Conf.AutorejoinBackupBinlog = !cluster.Conf.AutorejoinBackupBinlog
}

func (cluster *Cluster) SwitchRejoinSemisync() {
	cluster.Conf.AutorejoinSemisync = !cluster.Conf.AutorejoinSemisync
}
func (cluster *Cluster) SwitchRejoinFlashback() {
	cluster.Conf.AutorejoinFlashback = !cluster.Conf.AutorejoinFlashback
}

func (cluster *Cluster) SwitchRejoinPseudoGTID() {
	cluster.Conf.AutorejoinSlavePositionalHearbeat = !cluster.Conf.AutorejoinSlavePositionalHearbeat
}

func (cluster *Cluster) SwitchCheckReplicationFilters() {
	cluster.Conf.CheckReplFilter = !cluster.Conf.CheckReplFilter
}

func (cluster *Cluster) SwitchFailoverRestartUnsafe() {
	cluster.Conf.FailRestartUnsafe = !cluster.Conf.FailRestartUnsafe
}

func (cluster *Cluster) SwitchFailoverEventScheduler() {
	cluster.Conf.FailEventScheduler = !cluster.Conf.FailEventScheduler
}

func (cluster *Cluster) SwitchRejoinZFSFlashback() {
	cluster.Conf.AutorejoinZFSFlashback = !cluster.Conf.AutorejoinZFSFlashback
}

func (cluster *Cluster) SwitchBackup() {
	cluster.Conf.Backup = !cluster.Conf.Backup
}

func (cluster *Cluster) SwitchSchedulerBackupLogical() {
	cluster.Conf.SchedulerBackupLogical = !cluster.Conf.SchedulerBackupLogical
}

func (cluster *Cluster) SwitchSchedulerBackupPhysical() {
	cluster.Conf.SchedulerBackupPhysical = !cluster.Conf.SchedulerBackupPhysical
}

func (cluster *Cluster) SwitchSchedulerDatabaseLogs() {
	cluster.Conf.SchedulerDatabaseLogs = !cluster.Conf.SchedulerDatabaseLogs
}

func (cluster *Cluster) SwitchSchedulerDatabaseOptimize() {
	cluster.Conf.SchedulerDatabaseOptimize = !cluster.Conf.SchedulerDatabaseOptimize
}

func (cluster *Cluster) SwitchGraphiteEmbedded() {
	cluster.Conf.GraphiteEmbedded = !cluster.Conf.GraphiteEmbedded
}

func (cluster *Cluster) SwitchGraphiteMetrics() {
	cluster.Conf.GraphiteMetrics = !cluster.Conf.GraphiteMetrics
}

func (cluster *Cluster) SwitchFailoverEventStatus() {
	cluster.Conf.FailEventStatus = !cluster.Conf.FailEventStatus
}

func (cluster *Cluster) SwitchProxysqlBootstrap() {
	cluster.Conf.ProxysqlBootstrap = !cluster.Conf.ProxysqlBootstrap
}

func (cluster *Cluster) SwitchProxysqlCopyGrants() {
	cluster.Conf.ProxysqlCopyGrants = !cluster.Conf.ProxysqlCopyGrants
}

func (cluster *Cluster) SwitchMonitoringSchemaChange() {
	cluster.Conf.MonitorSchemaChange = !cluster.Conf.MonitorSchemaChange
}

func (cluster *Cluster) SwitchMonitoringScheduler() {
	cluster.Conf.MonitorScheduler = !cluster.Conf.MonitorScheduler
}

func (cluster *Cluster) SwitchMonitoringQueries() {
	cluster.Conf.MonitorQueries = !cluster.Conf.MonitorQueries
}

func (cluster *Cluster) SwitchTestMode() {
	cluster.Conf.Test = !cluster.Conf.Test
}

func (cluster *Cluster) SwitchTraffic() {
	cluster.SetTraffic(!cluster.GetTraffic())
}
