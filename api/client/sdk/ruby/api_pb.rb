# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
require 'google/api/annotations_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "openstorage.api.StorageResource" do
    optional :id, :string, 1
    optional :path, :string, 2
    optional :medium, :enum, 3, "openstorage.api.StorageMedium"
    optional :online, :bool, 4
    optional :iops, :uint64, 5
    optional :seq_write, :double, 6
    optional :seq_read, :double, 7
    optional :randRW, :double, 8
    optional :size, :uint64, 9
    optional :used, :uint64, 10
    optional :rotation_speed, :string, 11
    optional :last_scan, :message, 12, "google.protobuf.Timestamp"
    optional :metadata, :bool, 13
  end
  add_message "openstorage.api.StoragePool" do
    optional :ID, :int32, 1
    optional :Cos, :enum, 2, "openstorage.api.CosType"
    optional :Medium, :enum, 3, "openstorage.api.StorageMedium"
    optional :RaidLevel, :string, 4
    optional :TotalSize, :uint64, 7
    optional :Used, :uint64, 8
    map :labels, :string, :string, 9
  end
  add_message "openstorage.api.VolumeLocator" do
    optional :name, :string, 1
    map :volume_labels, :string, :string, 2
  end
  add_message "openstorage.api.Source" do
    optional :parent, :string, 1
    optional :seed, :string, 2
  end
  add_message "openstorage.api.Group" do
    optional :id, :string, 1
  end
  add_message "openstorage.api.VolumeSpec" do
    optional :ephemeral, :bool, 1
    optional :size, :uint64, 2
    optional :format, :enum, 3, "openstorage.api.FSType"
    optional :block_size, :int64, 4
    optional :ha_level, :int64, 5
    optional :cos, :enum, 6, "openstorage.api.CosType"
    optional :io_profile, :enum, 7, "openstorage.api.IoProfile"
    optional :dedupe, :bool, 8
    optional :snapshot_interval, :uint32, 9
    map :volume_labels, :string, :string, 10
    optional :shared, :bool, 11
    optional :replica_set, :message, 12, "openstorage.api.ReplicaSet"
    optional :aggregation_level, :uint32, 13
    optional :encrypted, :bool, 14
    optional :passphrase, :string, 15
    optional :snapshot_schedule, :string, 16
    optional :scale, :uint32, 17
    optional :sticky, :bool, 18
    optional :group, :message, 21, "openstorage.api.Group"
    optional :group_enforced, :bool, 22
    optional :compressed, :bool, 23
    optional :cascaded, :bool, 24
    optional :journal, :bool, 25
    optional :sharedv4, :bool, 26
  end
  add_message "openstorage.api.ReplicaSet" do
    repeated :nodes, :string, 1
  end
  add_message "openstorage.api.RuntimeStateMap" do
    map :runtime_state, :string, :string, 1
  end
  add_message "openstorage.api.Volume" do
    optional :id, :string, 1
    optional :source, :message, 2, "openstorage.api.Source"
    optional :group, :message, 3, "openstorage.api.Group"
    optional :readonly, :bool, 4
    optional :locator, :message, 5, "openstorage.api.VolumeLocator"
    optional :ctime, :message, 6, "google.protobuf.Timestamp"
    optional :spec, :message, 7, "openstorage.api.VolumeSpec"
    optional :usage, :uint64, 8
    optional :last_scan, :message, 9, "google.protobuf.Timestamp"
    optional :format, :enum, 10, "openstorage.api.FSType"
    optional :status, :enum, 11, "openstorage.api.VolumeStatus"
    optional :state, :enum, 12, "openstorage.api.VolumeState"
    optional :attached_on, :string, 13
    optional :attached_state, :enum, 14, "openstorage.api.AttachState"
    optional :device_path, :string, 15
    optional :secure_device_path, :string, 16
    repeated :attach_path, :string, 17
    map :attach_info, :string, :string, 18
    repeated :replica_sets, :message, 19, "openstorage.api.ReplicaSet"
    repeated :runtime_state, :message, 20, "openstorage.api.RuntimeStateMap"
    optional :error, :string, 21
    repeated :volume_consumers, :message, 22, "openstorage.api.VolumeConsumer"
  end
  add_message "openstorage.api.Stats" do
    optional :reads, :uint64, 1
    optional :read_ms, :uint64, 2
    optional :read_bytes, :uint64, 3
    optional :writes, :uint64, 4
    optional :write_ms, :uint64, 5
    optional :write_bytes, :uint64, 6
    optional :io_progress, :uint64, 7
    optional :io_ms, :uint64, 8
    optional :bytes_used, :uint64, 9
    optional :interval_ms, :uint64, 10
  end
  add_message "openstorage.api.Alert" do
    optional :id, :int64, 1
    optional :severity, :enum, 2, "openstorage.api.SeverityType"
    optional :alert_type, :int64, 3
    optional :message, :string, 4
    optional :timestamp, :message, 5, "google.protobuf.Timestamp"
    optional :resource_id, :string, 6
    optional :resource, :enum, 7, "openstorage.api.ResourceType"
    optional :cleared, :bool, 8
    optional :ttl, :uint64, 9
    optional :unique_tag, :string, 10
  end
  add_message "openstorage.api.Alerts" do
    repeated :alert, :message, 1, "openstorage.api.Alert"
  end
  add_message "openstorage.api.ObjectstoreInfo" do
    optional :uuid, :string, 1
    optional :volume_id, :string, 2
    optional :enabled, :bool, 3
    optional :status, :string, 4
    optional :action, :int64, 5
    optional :access_key, :string, 6
    optional :secret_key, :string, 7
    repeated :endpoints, :string, 8
    optional :current_endPoint, :string, 9
    optional :access_port, :int64, 10
    optional :region, :string, 11
  end
  add_message "openstorage.api.VolumeCreateRequest" do
    optional :locator, :message, 1, "openstorage.api.VolumeLocator"
    optional :source, :message, 2, "openstorage.api.Source"
    optional :spec, :message, 3, "openstorage.api.VolumeSpec"
  end
  add_message "openstorage.api.VolumeResponse" do
    optional :error, :string, 1
  end
  add_message "openstorage.api.VolumeCreateResponse" do
    optional :id, :string, 1
    optional :volume_response, :message, 2, "openstorage.api.VolumeResponse"
  end
  add_message "openstorage.api.VolumeStateAction" do
    optional :attach, :enum, 1, "openstorage.api.VolumeActionParam"
    optional :mount, :enum, 2, "openstorage.api.VolumeActionParam"
    optional :mount_path, :string, 3
    optional :device_path, :string, 4
  end
  add_message "openstorage.api.VolumeSetRequest" do
    optional :locator, :message, 1, "openstorage.api.VolumeLocator"
    optional :spec, :message, 2, "openstorage.api.VolumeSpec"
    optional :action, :message, 3, "openstorage.api.VolumeStateAction"
    map :options, :string, :string, 4
  end
  add_message "openstorage.api.VolumeSetResponse" do
    optional :volume, :message, 1, "openstorage.api.Volume"
    optional :volume_response, :message, 2, "openstorage.api.VolumeResponse"
  end
  add_message "openstorage.api.SnapCreateRequest" do
    optional :id, :string, 1
    optional :locator, :message, 2, "openstorage.api.VolumeLocator"
    optional :readonly, :bool, 3
  end
  add_message "openstorage.api.SnapCreateResponse" do
    optional :volume_create_response, :message, 1, "openstorage.api.VolumeCreateResponse"
  end
  add_message "openstorage.api.VolumeInfo" do
    optional :volume_id, :string, 1
    optional :path, :string, 2
    optional :storage, :message, 3, "openstorage.api.VolumeSpec"
  end
  add_message "openstorage.api.VolumeConsumer" do
    optional :name, :string, 1
    optional :namespace, :string, 2
    optional :type, :string, 3
    optional :node_id, :string, 4
    optional :owner_name, :string, 5
    optional :owner_type, :string, 6
  end
  add_message "openstorage.api.GraphDriverChanges" do
    optional :path, :string, 1
    optional :kind, :enum, 2, "openstorage.api.GraphDriverChangeType"
  end
  add_message "openstorage.api.ClusterResponse" do
    optional :error, :string, 1
  end
  add_message "openstorage.api.ActiveRequest" do
    map :ReqestKV, :int64, :string, 1
  end
  add_message "openstorage.api.ActiveRequests" do
    optional :RequestCount, :int64, 1
    repeated :ActiveRequest, :message, 2, "openstorage.api.ActiveRequest"
  end
  add_message "openstorage.api.GroupSnapCreateRequest" do
    optional :id, :string, 1
    map :Labels, :string, :string, 2
  end
  add_message "openstorage.api.GroupSnapCreateResponse" do
    map :snapshots, :string, :message, 1, "openstorage.api.SnapCreateResponse"
    optional :error, :string, 2
  end
  add_message "openstorage.api.StorageNode" do
    optional :id, :string, 1
    optional :cpu, :double, 2
    optional :mem_total, :uint64, 3
    optional :mem_used, :uint64, 4
    optional :mem_free, :uint64, 5
    optional :avg_load, :int64, 6
    optional :status, :enum, 7, "openstorage.api.Status"
    map :disks, :string, :message, 9, "openstorage.api.StorageResource"
    repeated :pools, :message, 10, "openstorage.api.StoragePool"
    optional :mgmt_ip, :string, 11
    optional :data_ip, :string, 12
    optional :hostname, :string, 15
    map :node_labels, :string, :string, 16
  end
  add_message "openstorage.api.StorageCluster" do
    optional :status, :enum, 1, "openstorage.api.Status"
    optional :id, :string, 2
    optional :node_id, :string, 3
    repeated :nodes, :message, 4, "openstorage.api.StorageNode"
  end
  add_message "openstorage.api.VolumeMountRequest" do
    optional :volume_id, :string, 1
    optional :mount_path, :string, 2
    map :options, :string, :string, 3
  end
  add_message "openstorage.api.VolumeMountResponse" do
  end
  add_message "openstorage.api.VolumeUnmountRequest" do
    optional :volume_id, :string, 1
    optional :mount_path, :string, 2
    map :options, :string, :string, 3
  end
  add_message "openstorage.api.VolumeUnmountResponse" do
  end
  add_message "openstorage.api.VolumeAttachRequest" do
    optional :volume_id, :string, 1
    map :options, :string, :string, 2
  end
  add_message "openstorage.api.VolumeAttachResponse" do
    optional :device_path, :string, 1
  end
  add_message "openstorage.api.VolumeDetachRequest" do
    optional :volume_id, :string, 1
  end
  add_message "openstorage.api.VolumeDetachResponse" do
  end
  add_message "openstorage.api.OpenStorageVolumeCreateRequest" do
    optional :name, :string, 1
    optional :spec, :message, 2, "openstorage.api.VolumeSpec"
  end
  add_message "openstorage.api.OpenStorageVolumeCreateResponse" do
    optional :volume_id, :string, 1
  end
  add_message "openstorage.api.VolumeCreateFromVolumeIDRequest" do
    optional :name, :string, 1
    optional :parent_id, :string, 2
    optional :spec, :message, 3, "openstorage.api.VolumeSpec"
  end
  add_message "openstorage.api.VolumeCreateFromVolumeIDResponse" do
    optional :volume_id, :string, 1
  end
  add_message "openstorage.api.VolumeDeleteRequest" do
    optional :volume_id, :string, 1
  end
  add_message "openstorage.api.VolumeDeleteResponse" do
  end
  add_message "openstorage.api.VolumeInspectRequest" do
    optional :volume_id, :string, 1
  end
  add_message "openstorage.api.VolumeInspectResponse" do
    optional :volume, :message, 1, "openstorage.api.Volume"
  end
  add_message "openstorage.api.VolumeEnumerateRequest" do
    optional :locator, :message, 1, "openstorage.api.VolumeLocator"
  end
  add_message "openstorage.api.VolumeEnumerateResponse" do
    repeated :volumes, :message, 1, "openstorage.api.Volume"
  end
  add_message "openstorage.api.VolumeSnapshotCreateRequest" do
    optional :volume_id, :string, 1
    map :labels, :string, :string, 2
  end
  add_message "openstorage.api.VolumeSnapshotCreateResponse" do
    optional :snapshot_id, :string, 1
  end
  add_message "openstorage.api.VolumeSnapshotRestoreRequest" do
    optional :volume_id, :string, 1
    optional :snapshot_id, :string, 2
  end
  add_message "openstorage.api.VolumeSnapshotRestoreResponse" do
  end
  add_message "openstorage.api.VolumeSnapshotEnumerateRequest" do
    optional :volume_id, :string, 1
    map :labels, :string, :string, 2
  end
  add_message "openstorage.api.VolumeSnapshotEnumerateResponse" do
    repeated :snapshots, :message, 1, "openstorage.api.Volume"
  end
  add_message "openstorage.api.ClusterEnumerateRequest" do
  end
  add_message "openstorage.api.ClusterEnumerateResponse" do
    optional :cluster, :message, 1, "openstorage.api.StorageCluster"
  end
  add_message "openstorage.api.ClusterInspectRequest" do
    optional :node_id, :string, 1
  end
  add_message "openstorage.api.ClusterInspectResponse" do
    optional :node, :message, 1, "openstorage.api.StorageNode"
  end
  add_message "openstorage.api.ClusterAlertEnumerateRequest" do
    optional :time_start, :message, 1, "google.protobuf.Timestamp"
    optional :time_end, :message, 2, "google.protobuf.Timestamp"
    optional :resource, :enum, 3, "openstorage.api.ResourceType"
  end
  add_message "openstorage.api.ClusterAlertEnumerateResponse" do
    optional :alerts, :message, 1, "openstorage.api.Alerts"
  end
  add_message "openstorage.api.ClusterAlertClearRequest" do
    optional :resource, :enum, 1, "openstorage.api.ResourceType"
    optional :alert_id, :int64, 2
  end
  add_message "openstorage.api.ClusterAlertClearResponse" do
  end
  add_message "openstorage.api.ClusterAlertEraseRequest" do
    optional :resource, :enum, 1, "openstorage.api.ResourceType"
    optional :alert_id, :int64, 2
  end
  add_message "openstorage.api.ClusterAlertEraseResponse" do
  end
  add_message "openstorage.api.CloudBackupCreateRequest" do
    optional :volume_id, :string, 1
    optional :credential_uuid, :string, 2
    optional :full, :bool, 3
  end
  add_message "openstorage.api.CloudBackupCreateResponse" do
  end
  add_message "openstorage.api.CloudBackupRestoreRequest" do
    optional :id, :string, 1
    optional :restore_volume_name, :string, 2
    optional :credential_uuid, :string, 3
    optional :node_id, :string, 4
  end
  add_message "openstorage.api.CloudBackupRestoreResponse" do
    optional :restore_volume_id, :string, 1
  end
  add_message "openstorage.api.CloudBackupGenericRequest" do
    optional :src_volume_id, :string, 1
    optional :cluster_id, :string, 2
    optional :credential_uuid, :string, 3
    optional :all, :bool, 4
  end
  add_message "openstorage.api.CloudBackupInfo" do
    optional :id, :string, 1
    optional :src_volume_id, :string, 2
    optional :src_volume_name, :string, 3
    optional :timestamp, :message, 4, "google.protobuf.Timestamp"
    map :metadata, :string, :string, 5
    optional :status, :string, 6
  end
  add_message "openstorage.api.CloudBackupEnumerateRequest" do
    optional :src_volume_id, :string, 1
    optional :cluster_id, :string, 2
    optional :credential_uuid, :string, 3
    optional :all, :bool, 4
  end
  add_message "openstorage.api.CloudBackupEnumerateResponse" do
    repeated :backups, :message, 1, "openstorage.api.CloudBackupInfo"
  end
  add_message "openstorage.api.CloudBackupDeleteRequest" do
    optional :id, :string, 1
    optional :credential_uuid, :string, 2
    optional :force, :bool, 3
  end
  add_message "openstorage.api.CloudBackupDeleteAllRequest" do
    optional :src_volume_id, :string, 1
    optional :cluster_id, :string, 2
    optional :credential_uuid, :string, 3
    optional :all, :bool, 4
  end
  add_message "openstorage.api.CloudBackupStatusRequest" do
    optional :src_volume_id, :string, 1
    optional :local, :bool, 2
  end
  add_message "openstorage.api.CloudBackupStatus" do
    optional :id, :string, 1
    optional :optype, :enum, 2, "openstorage.api.CloudBackupOpType"
    optional :status, :enum, 3, "openstorage.api.CloudBackupStatusType"
    optional :bytes_done, :uint64, 4
    optional :start_time, :message, 5, "google.protobuf.Timestamp"
    optional :comleted_time, :message, 6, "google.protobuf.Timestamp"
    optional :node_id, :string, 7
  end
  add_message "openstorage.api.CloudBackupStatusResponse" do
    map :statuses, :string, :message, 1, "openstorage.api.CloudBackupStatus"
  end
  add_message "openstorage.api.CloudBackupCatalogRequest" do
    optional :id, :string, 1
    optional :credential_uuid, :string, 2
  end
  add_message "openstorage.api.CloudBackupCatalogResponse" do
    repeated :contents, :string, 1
  end
  add_message "openstorage.api.CloudBackupHistoryRequest" do
    optional :src_volume_id, :string, 1
  end
  add_message "openstorage.api.CloudBackupHistoryItem" do
    optional :src_volume_id, :string, 1
    optional :timestamp, :message, 2, "google.protobuf.Timestamp"
    optional :status, :string, 3
  end
  add_message "openstorage.api.CloudBackupHistoryResponse" do
    repeated :history_list, :message, 1, "openstorage.api.CloudBackupHistoryItem"
  end
  add_message "openstorage.api.CloudBackupStateChangeRequest" do
    optional :src_volume_id, :string, 1
    optional :requested_state, :string, 2
  end
  add_message "openstorage.api.CloudBackupScheduleInfo" do
    optional :src_volume_id, :string, 1
    optional :credential_uuid, :string, 2
    optional :schedule, :string, 3
    optional :max_backups, :uint32, 4
  end
  add_message "openstorage.api.CloudBackupSchedCreateRequest" do
    optional :src_volume_id, :string, 1
    optional :credential_uuid, :string, 2
    optional :schedule, :string, 3
    optional :max_backups, :uint32, 4
  end
  add_message "openstorage.api.CloudBackupSchedCreateResponse" do
    optional :uuid, :string, 1
  end
  add_message "openstorage.api.CloudBackupSchedDeleteRequest" do
    optional :uuid, :string, 1
  end
  add_message "openstorage.api.CloudBackupSchedEnumerateResponse" do
    map :schedules, :string, :message, 1, "openstorage.api.CloudBackupScheduleInfo"
  end
  add_enum "openstorage.api.Status" do
    value :STATUS_NONE, 0
    value :STATUS_INIT, 1
    value :STATUS_OK, 2
    value :STATUS_OFFLINE, 3
    value :STATUS_ERROR, 4
    value :STATUS_NOT_IN_QUORUM, 5
    value :STATUS_DECOMMISSION, 6
    value :STATUS_MAINTENANCE, 7
    value :STATUS_STORAGE_DOWN, 8
    value :STATUS_STORAGE_DEGRADED, 9
    value :STATUS_NEEDS_REBOOT, 10
    value :STATUS_STORAGE_REBALANCE, 11
    value :STATUS_STORAGE_DRIVE_REPLACE, 12
    value :STATUS_NOT_IN_QUORUM_NO_STORAGE, 13
    value :STATUS_MAX, 14
  end
  add_enum "openstorage.api.DriverType" do
    value :DRIVER_TYPE_NONE, 0
    value :DRIVER_TYPE_FILE, 1
    value :DRIVER_TYPE_BLOCK, 2
    value :DRIVER_TYPE_OBJECT, 3
    value :DRIVER_TYPE_CLUSTERED, 4
    value :DRIVER_TYPE_GRAPH, 5
  end
  add_enum "openstorage.api.FSType" do
    value :FS_TYPE_NONE, 0
    value :FS_TYPE_BTRFS, 1
    value :FS_TYPE_EXT4, 2
    value :FS_TYPE_FUSE, 3
    value :FS_TYPE_NFS, 4
    value :FS_TYPE_VFS, 5
    value :FS_TYPE_XFS, 6
    value :FS_TYPE_ZFS, 7
  end
  add_enum "openstorage.api.GraphDriverChangeType" do
    value :GRAPH_DRIVER_CHANGE_TYPE_NONE, 0
    value :GRAPH_DRIVER_CHANGE_TYPE_MODIFIED, 1
    value :GRAPH_DRIVER_CHANGE_TYPE_ADDED, 2
    value :GRAPH_DRIVER_CHANGE_TYPE_DELETED, 3
  end
  add_enum "openstorage.api.SeverityType" do
    value :SEVERITY_TYPE_NONE, 0
    value :SEVERITY_TYPE_ALARM, 1
    value :SEVERITY_TYPE_WARNING, 2
    value :SEVERITY_TYPE_NOTIFY, 3
  end
  add_enum "openstorage.api.ResourceType" do
    value :RESOURCE_TYPE_NONE, 0
    value :RESOURCE_TYPE_VOLUME, 1
    value :RESOURCE_TYPE_NODE, 2
    value :RESOURCE_TYPE_CLUSTER, 3
    value :RESOURCE_TYPE_DRIVE, 4
  end
  add_enum "openstorage.api.AlertActionType" do
    value :ALERT_ACTION_TYPE_NONE, 0
    value :ALERT_ACTION_TYPE_DELETE, 1
    value :ALERT_ACTION_TYPE_CREATE, 2
    value :ALERT_ACTION_TYPE_UPDATE, 3
  end
  add_enum "openstorage.api.VolumeActionParam" do
    value :VOLUME_ACTION_PARAM_NONE, 0
    value :VOLUME_ACTION_PARAM_OFF, 1
    value :VOLUME_ACTION_PARAM_ON, 2
  end
  add_enum "openstorage.api.CosType" do
    value :NONE, 0
    value :LOW, 1
    value :MEDIUM, 2
    value :HIGH, 3
  end
  add_enum "openstorage.api.IoProfile" do
    value :IO_PROFILE_SEQUENTIAL, 0
    value :IO_PROFILE_RANDOM, 1
    value :IO_PROFILE_DB, 2
    value :IO_PROFILE_DB_REMOTE, 3
    value :IO_PROFILE_CMS, 4
  end
  add_enum "openstorage.api.VolumeState" do
    value :VOLUME_STATE_NONE, 0
    value :VOLUME_STATE_PENDING, 1
    value :VOLUME_STATE_AVAILABLE, 2
    value :VOLUME_STATE_ATTACHED, 3
    value :VOLUME_STATE_DETACHED, 4
    value :VOLUME_STATE_DETATCHING, 5
    value :VOLUME_STATE_ERROR, 6
    value :VOLUME_STATE_DELETED, 7
    value :VOLUME_STATE_TRY_DETACHING, 8
    value :VOLUME_STATE_RESTORE, 9
  end
  add_enum "openstorage.api.VolumeStatus" do
    value :VOLUME_STATUS_NONE, 0
    value :VOLUME_STATUS_NOT_PRESENT, 1
    value :VOLUME_STATUS_UP, 2
    value :VOLUME_STATUS_DOWN, 3
    value :VOLUME_STATUS_DEGRADED, 4
  end
  add_enum "openstorage.api.StorageMedium" do
    value :STORAGE_MEDIUM_MAGNETIC, 0
    value :STORAGE_MEDIUM_SSD, 1
    value :STORAGE_MEDIUM_NVME, 2
  end
  add_enum "openstorage.api.ClusterNotify" do
    value :CLUSTER_NOTIFY_DOWN, 0
  end
  add_enum "openstorage.api.AttachState" do
    value :ATTACH_STATE_EXTERNAL, 0
    value :ATTACH_STATE_INTERNAL, 1
    value :ATTACH_STATE_INTERNAL_SWITCH, 2
  end
  add_enum "openstorage.api.OperationFlags" do
    value :OP_FLAGS_UNKNOWN, 0
    value :OP_FLAGS_NONE, 1
    value :OP_FLAGS_DETACH_FORCE, 2
  end
  add_enum "openstorage.api.CloudBackupOpType" do
    value :CloudBackupOp, 0
    value :CloudRestoreOp, 1
  end
  add_enum "openstorage.api.CloudBackupStatusType" do
    value :CloudBackupStatusNotStarted, 0
    value :CloudBackupStatusDone, 1
    value :CloudBackupStatusAborted, 2
    value :CloudBackupStatusPaused, 3
    value :CloudBackupStatusStopped, 4
    value :CloudBackupStatusActive, 5
    value :CloudBackupStatusFailed, 6
  end
end

module Openstorage
  module Api
    StorageResource = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.StorageResource").msgclass
    StoragePool = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.StoragePool").msgclass
    VolumeLocator = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeLocator").msgclass
    Source = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Source").msgclass
    Group = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Group").msgclass
    VolumeSpec = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSpec").msgclass
    ReplicaSet = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ReplicaSet").msgclass
    RuntimeStateMap = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.RuntimeStateMap").msgclass
    Volume = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Volume").msgclass
    Stats = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Stats").msgclass
    Alert = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Alert").msgclass
    Alerts = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Alerts").msgclass
    ObjectstoreInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ObjectstoreInfo").msgclass
    VolumeCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeCreateRequest").msgclass
    VolumeResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeResponse").msgclass
    VolumeCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeCreateResponse").msgclass
    VolumeStateAction = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeStateAction").msgclass
    VolumeSetRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSetRequest").msgclass
    VolumeSetResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSetResponse").msgclass
    SnapCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.SnapCreateRequest").msgclass
    SnapCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.SnapCreateResponse").msgclass
    VolumeInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeInfo").msgclass
    VolumeConsumer = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeConsumer").msgclass
    GraphDriverChanges = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.GraphDriverChanges").msgclass
    ClusterResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterResponse").msgclass
    ActiveRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ActiveRequest").msgclass
    ActiveRequests = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ActiveRequests").msgclass
    GroupSnapCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.GroupSnapCreateRequest").msgclass
    GroupSnapCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.GroupSnapCreateResponse").msgclass
    StorageNode = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.StorageNode").msgclass
    StorageCluster = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.StorageCluster").msgclass
    VolumeMountRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeMountRequest").msgclass
    VolumeMountResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeMountResponse").msgclass
    VolumeUnmountRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeUnmountRequest").msgclass
    VolumeUnmountResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeUnmountResponse").msgclass
    VolumeAttachRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeAttachRequest").msgclass
    VolumeAttachResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeAttachResponse").msgclass
    VolumeDetachRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeDetachRequest").msgclass
    VolumeDetachResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeDetachResponse").msgclass
    OpenStorageVolumeCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.OpenStorageVolumeCreateRequest").msgclass
    OpenStorageVolumeCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.OpenStorageVolumeCreateResponse").msgclass
    VolumeCreateFromVolumeIDRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeCreateFromVolumeIDRequest").msgclass
    VolumeCreateFromVolumeIDResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeCreateFromVolumeIDResponse").msgclass
    VolumeDeleteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeDeleteRequest").msgclass
    VolumeDeleteResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeDeleteResponse").msgclass
    VolumeInspectRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeInspectRequest").msgclass
    VolumeInspectResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeInspectResponse").msgclass
    VolumeEnumerateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeEnumerateRequest").msgclass
    VolumeEnumerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeEnumerateResponse").msgclass
    VolumeSnapshotCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSnapshotCreateRequest").msgclass
    VolumeSnapshotCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSnapshotCreateResponse").msgclass
    VolumeSnapshotRestoreRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSnapshotRestoreRequest").msgclass
    VolumeSnapshotRestoreResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSnapshotRestoreResponse").msgclass
    VolumeSnapshotEnumerateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSnapshotEnumerateRequest").msgclass
    VolumeSnapshotEnumerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeSnapshotEnumerateResponse").msgclass
    ClusterEnumerateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterEnumerateRequest").msgclass
    ClusterEnumerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterEnumerateResponse").msgclass
    ClusterInspectRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterInspectRequest").msgclass
    ClusterInspectResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterInspectResponse").msgclass
    ClusterAlertEnumerateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterAlertEnumerateRequest").msgclass
    ClusterAlertEnumerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterAlertEnumerateResponse").msgclass
    ClusterAlertClearRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterAlertClearRequest").msgclass
    ClusterAlertClearResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterAlertClearResponse").msgclass
    ClusterAlertEraseRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterAlertEraseRequest").msgclass
    ClusterAlertEraseResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterAlertEraseResponse").msgclass
    CloudBackupCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupCreateRequest").msgclass
    CloudBackupCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupCreateResponse").msgclass
    CloudBackupRestoreRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupRestoreRequest").msgclass
    CloudBackupRestoreResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupRestoreResponse").msgclass
    CloudBackupGenericRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupGenericRequest").msgclass
    CloudBackupInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupInfo").msgclass
    CloudBackupEnumerateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupEnumerateRequest").msgclass
    CloudBackupEnumerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupEnumerateResponse").msgclass
    CloudBackupDeleteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupDeleteRequest").msgclass
    CloudBackupDeleteAllRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupDeleteAllRequest").msgclass
    CloudBackupStatusRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupStatusRequest").msgclass
    CloudBackupStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupStatus").msgclass
    CloudBackupStatusResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupStatusResponse").msgclass
    CloudBackupCatalogRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupCatalogRequest").msgclass
    CloudBackupCatalogResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupCatalogResponse").msgclass
    CloudBackupHistoryRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupHistoryRequest").msgclass
    CloudBackupHistoryItem = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupHistoryItem").msgclass
    CloudBackupHistoryResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupHistoryResponse").msgclass
    CloudBackupStateChangeRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupStateChangeRequest").msgclass
    CloudBackupScheduleInfo = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupScheduleInfo").msgclass
    CloudBackupSchedCreateRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupSchedCreateRequest").msgclass
    CloudBackupSchedCreateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupSchedCreateResponse").msgclass
    CloudBackupSchedDeleteRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupSchedDeleteRequest").msgclass
    CloudBackupSchedEnumerateResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupSchedEnumerateResponse").msgclass
    Status = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.Status").enummodule
    DriverType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.DriverType").enummodule
    FSType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.FSType").enummodule
    GraphDriverChangeType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.GraphDriverChangeType").enummodule
    SeverityType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.SeverityType").enummodule
    ResourceType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ResourceType").enummodule
    AlertActionType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.AlertActionType").enummodule
    VolumeActionParam = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeActionParam").enummodule
    CosType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CosType").enummodule
    IoProfile = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.IoProfile").enummodule
    VolumeState = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeState").enummodule
    VolumeStatus = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.VolumeStatus").enummodule
    StorageMedium = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.StorageMedium").enummodule
    ClusterNotify = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.ClusterNotify").enummodule
    AttachState = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.AttachState").enummodule
    OperationFlags = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.OperationFlags").enummodule
    CloudBackupOpType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupOpType").enummodule
    CloudBackupStatusType = Google::Protobuf::DescriptorPool.generated_pool.lookup("openstorage.api.CloudBackupStatusType").enummodule
  end
end
