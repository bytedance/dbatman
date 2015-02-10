// Copyright

%{

package sql

import (
)

%}


%union {
    Bytes []byte
    statement IStatement
    Table *TableInfo
    empty struct{}
}

/*
   Comments for TOKENS.
   For each token, please include in the same line a comment that contains
   the following tags:
   SQL-2003-R : Reserved keyword as per SQL-2003
   SQL-2003-N : Non Reserved keyword as per SQL-2003
   SQL-1999-R : Reserved keyword as per SQL-1999
   SQL-1999-N : Non Reserved keyword as per SQL-1999
   MYSQL      : MySQL extention (unspecified)
   MYSQL-FUNC : MySQL extention, function
   INTERNAL   : Not a real token, lex optimization
   OPERATOR   : SQL operator
   FUTURE-USE : Reserved for futur use

   This makes the code grep-able, and helps maintenance.
*/

%token  ABORT_SYM                     /* INTERNAL (used in lex) */
%token<Bytes>  ACCESSIBLE_SYM
%token<Bytes>  ACTION                        /* SQL-2003-N */
%token<Bytes>  ADD                           /* SQL-2003-R */
%token<Bytes>  ADDDATE_SYM                   /* MYSQL-FUNC */
%token<Bytes>  AFTER_SYM                     /* SQL-2003-N */
%token<Bytes>  AGAINST
%token<Bytes>  AGGREGATE_SYM
%token<Bytes>  ALGORITHM_SYM
%token<Bytes>  ALL                           /* SQL-2003-R */
%token<Bytes>  ALTER                         /* SQL-2003-R */
%token<Bytes>  ANALYSE_SYM
%token<Bytes>  ANALYZE_SYM
%token<Bytes>  AND_AND_SYM                   /* OPERATOR */
%token<Bytes>  AND_SYM                       /* SQL-2003-R */
%token<Bytes>  ANY_SYM                       /* SQL-2003-R */
%token<Bytes>  AS                            /* SQL-2003-R */
%token<Bytes>  ASC                           /* SQL-2003-N */
%token<Bytes>  ASCII_SYM                     /* MYSQL-FUNC */
%token<Bytes>  ASENSITIVE_SYM                /* FUTURE-USE */
%token<Bytes>  AT_SYM                        /* SQL-2003-R */
%token<Bytes>  AUTOEXTEND_SIZE_SYM
%token<Bytes>  AUTO_INC
%token<Bytes>  AVG_ROW_LENGTH
%token<Bytes>  AVG_SYM                       /* SQL-2003-N */
%token<Bytes>  BACKUP_SYM
%token<Bytes>  BEFORE_SYM                    /* SQL-2003-N */
%token<Bytes>  BEGIN_SYM                     /* SQL-2003-R */
%token<Bytes>  BETWEEN_SYM                   /* SQL-2003-R */
%token<Bytes>  BIGINT                        /* SQL-2003-R */
%token<Bytes>  BINARY                        /* SQL-2003-R */
%token<Bytes>  BINLOG_SYM
%token<Bytes>  BIN_NUM
%token<Bytes>  BIT_AND                       /* MYSQL-FUNC */
%token<Bytes>  BIT_OR                        /* MYSQL-FUNC */
%token<Bytes>  BIT_SYM                       /* MYSQL-FUNC */
%token<Bytes>  BIT_XOR                       /* MYSQL-FUNC */
%token<Bytes>  BLOB_SYM                      /* SQL-2003-R */
%token<Bytes>  BLOCK_SYM
%token<Bytes>  BOOLEAN_SYM                   /* SQL-2003-R */
%token<Bytes>  BOOL_SYM
%token<Bytes>  BOTH                          /* SQL-2003-R */
%token<Bytes>  BTREE_SYM
%token<Bytes>  BY                            /* SQL-2003-R */
%token<Bytes>  BYTE_SYM
%token<Bytes>  CACHE_SYM
%token<Bytes>  CALL_SYM                      /* SQL-2003-R */
%token<Bytes>  CASCADE                       /* SQL-2003-N */
%token<Bytes>  CASCADED                      /* SQL-2003-R */
%token<Bytes>  CASE_SYM                      /* SQL-2003-R */
%token<Bytes>  CAST_SYM                      /* SQL-2003-R */
%token<Bytes>  CATALOG_NAME_SYM              /* SQL-2003-N */
%token<Bytes>  CHAIN_SYM                     /* SQL-2003-N */
%token<Bytes>  CHANGE
%token<Bytes>  CHANGED
%token<Bytes>  CHARSET
%token<Bytes>  CHAR_SYM                      /* SQL-2003-R */
%token<Bytes>  CHECKSUM_SYM
%token<Bytes>  CHECK_SYM                     /* SQL-2003-R */
%token<Bytes>  CIPHER_SYM
%token<Bytes>  CLASS_ORIGIN_SYM              /* SQL-2003-N */
%token<Bytes>  CLIENT_SYM
%token<Bytes>  CLOSE_SYM                     /* SQL-2003-R */
%token<Bytes>  COALESCE                      /* SQL-2003-N */
%token<Bytes>  CODE_SYM
%token<Bytes>  COLLATE_SYM                   /* SQL-2003-R */
%token<Bytes>  COLLATION_SYM                 /* SQL-2003-N */
%token<Bytes>  COLUMNS
%token<Bytes>  COLUMN_SYM                    /* SQL-2003-R */
%token<Bytes>  COLUMN_FORMAT_SYM
%token<Bytes>  COLUMN_NAME_SYM               /* SQL-2003-N */
%token<Bytes>  COMMENT_SYM
%token<Bytes>  COMMITTED_SYM                 /* SQL-2003-N */
%token<Bytes>  COMMIT_SYM                    /* SQL-2003-R */
%token<Bytes>  COMPACT_SYM
%token<Bytes>  COMPLETION_SYM
%token<Bytes>  COMPRESSED_SYM
%token<Bytes>  CONCURRENT
%token<Bytes>  CONDITION_SYM                 /* SQL-2003-R, SQL-2008-R */
%token<Bytes>  CONNECTION_SYM
%token<Bytes>  CONSISTENT_SYM
%token<Bytes>  CONSTRAINT                    /* SQL-2003-R */
%token<Bytes>  CONSTRAINT_CATALOG_SYM        /* SQL-2003-N */
%token<Bytes>  CONSTRAINT_NAME_SYM           /* SQL-2003-N */
%token<Bytes>  CONSTRAINT_SCHEMA_SYM         /* SQL-2003-N */
%token<Bytes>  CONTAINS_SYM                  /* SQL-2003-N */
%token<Bytes>  CONTEXT_SYM
%token<Bytes>  CONTINUE_SYM                  /* SQL-2003-R */
%token<Bytes>  CONVERT_SYM                   /* SQL-2003-N */
%token<Bytes>  COUNT_SYM                     /* SQL-2003-N */
%token<Bytes>  CPU_SYM
%token<Bytes>  CREATE                        /* SQL-2003-R */
%token<Bytes>  CROSS                         /* SQL-2003-R */
%token<Bytes>  CUBE_SYM                      /* SQL-2003-R */
%token<Bytes>  CURDATE                       /* MYSQL-FUNC */
%token<Bytes>  CURRENT_SYM                   /* SQL-2003-R */
%token<Bytes>  CURRENT_USER                  /* SQL-2003-R */
%token<Bytes>  CURSOR_SYM                    /* SQL-2003-R */
%token<Bytes>  CURSOR_NAME_SYM               /* SQL-2003-N */
%token<Bytes>  CURTIME                       /* MYSQL-FUNC */
%token<Bytes>  DATABASE
%token<Bytes>  DATABASES
%token<Bytes>  DATAFILE_SYM
%token<Bytes>  DATA_SYM                      /* SQL-2003-N */
%token<Bytes>  DATETIME
%token<Bytes>  DATE_ADD_INTERVAL             /* MYSQL-FUNC */
%token<Bytes>  DATE_SUB_INTERVAL             /* MYSQL-FUNC */
%token<Bytes>  DATE_SYM                      /* SQL-2003-R */
%token<Bytes>  DAY_HOUR_SYM
%token<Bytes>  DAY_MICROSECOND_SYM
%token<Bytes>  DAY_MINUTE_SYM
%token<Bytes>  DAY_SECOND_SYM
%token<Bytes>  DAY_SYM                       /* SQL-2003-R */
%token<Bytes>  DEALLOCATE_SYM                /* SQL-2003-R */
%token<Bytes>  DECIMAL_NUM
%token<Bytes>  DECIMAL_SYM                   /* SQL-2003-R */
%token<Bytes>  DECLARE_SYM                   /* SQL-2003-R */
%token<Bytes>  DEFAULT                       /* SQL-2003-R */
%token<Bytes>  DEFAULT_AUTH_SYM              /* INTERNAL */
%token<Bytes>  DEFINER_SYM
%token<Bytes>  DELAYED_SYM
%token<Bytes>  DELAY_KEY_WRITE_SYM
%token<Bytes>  DELETE_SYM                    /* SQL-2003-R */
%token<Bytes>  DESC                          /* SQL-2003-N */
%token<Bytes>  DESCRIBE                      /* SQL-2003-R */
%token<Bytes>  DES_KEY_FILE
%token<Bytes>  DETERMINISTIC_SYM             /* SQL-2003-R */
%token<Bytes>  DIAGNOSTICS_SYM               /* SQL-2003-N */
%token<Bytes>  DIRECTORY_SYM
%token<Bytes>  DISABLE_SYM
%token<Bytes>  DISCARD
%token<Bytes>  DISK_SYM
%token<Bytes>  DISTINCT                      /* SQL-2003-R */
%token<Bytes>  DIV_SYM
%token<Bytes>  DOUBLE_SYM                    /* SQL-2003-R */
%token<Bytes>  DO_SYM
%token<Bytes>  DROP                          /* SQL-2003-R */
%token<Bytes>  DUAL_SYM
%token<Bytes>  DUMPFILE
%token<Bytes>  DUPLICATE_SYM
%token<Bytes>  DYNAMIC_SYM                   /* SQL-2003-R */
%token<Bytes>  EACH_SYM                      /* SQL-2003-R */
%token<Bytes>  ELSE                          /* SQL-2003-R */
%token<Bytes>  ELSEIF_SYM
%token<Bytes>  ENABLE_SYM
%token<Bytes>  ENCLOSED
%token<Bytes>  END                           /* SQL-2003-R */
%token<Bytes>  ENDS_SYM
%token<Bytes>  END_OF_INPUT                  /* INTERNAL */
%token<Bytes>  ENGINES_SYM
%token<Bytes>  ENGINE_SYM
%token<Bytes>  ENUM
%token<Bytes>  EQ                            /* OPERATOR */
%token<Bytes>  EQUAL_SYM                     /* OPERATOR */
%token<Bytes>  ERROR_SYM
%token<Bytes>  ERRORS
%token<Bytes>  ESCAPED
%token<Bytes>  ESCAPE_SYM                    /* SQL-2003-R */
%token<Bytes>  EVENTS_SYM
%token<Bytes>  EVENT_SYM
%token<Bytes>  EVERY_SYM                     /* SQL-2003-N */
%token<Bytes>  EXCHANGE_SYM
%token<Bytes>  EXECUTE_SYM                   /* SQL-2003-R */
%token<Bytes>  EXISTS                        /* SQL-2003-R */
%token<Bytes>  EXIT_SYM
%token<Bytes>  EXPANSION_SYM
%token<Bytes>  EXPIRE_SYM
%token<Bytes>  EXPORT_SYM
%token<Bytes>  EXTENDED_SYM
%token<Bytes>  EXTENT_SIZE_SYM
%token<Bytes>  EXTRACT_SYM                   /* SQL-2003-N */
%token<Bytes>  FALSE_SYM                     /* SQL-2003-R */
%token<Bytes>  FAST_SYM
%token<Bytes>  FAULTS_SYM
%token<Bytes>  FETCH_SYM                     /* SQL-2003-R */
%token<Bytes>  FILE_SYM
%token<Bytes>  FIRST_SYM                     /* SQL-2003-N */
%token<Bytes>  FIXED_SYM
%token<Bytes>  FLOAT_NUM
%token<Bytes>  FLOAT_SYM                     /* SQL-2003-R */
%token<Bytes>  FLUSH_SYM
%token<Bytes>  FORCE_SYM
%token<Bytes>  FOREIGN                       /* SQL-2003-R */
%token<Bytes>  FOR_SYM                       /* SQL-2003-R */
%token<Bytes>  FORMAT_SYM
%token<Bytes>  FOUND_SYM                     /* SQL-2003-R */
%token<Bytes>  FROM
%token<Bytes>  FULL                          /* SQL-2003-R */
%token<Bytes>  FULLTEXT_SYM
%token<Bytes>  FUNCTION_SYM                  /* SQL-2003-R */
%token<Bytes>  GE
%token<Bytes>  GENERAL
%token<Bytes>  GEOMETRYCOLLECTION
%token<Bytes>  GEOMETRY_SYM
%token<Bytes>  GET_FORMAT                    /* MYSQL-FUNC */
%token<Bytes>  GET_SYM                       /* SQL-2003-R */
%token<Bytes>  GLOBAL_SYM                    /* SQL-2003-R */
%token<Bytes>  GRANT                         /* SQL-2003-R */
%token<Bytes>  GRANTS
%token<Bytes>  GROUP_SYM                     /* SQL-2003-R */
%token<Bytes>  GROUP_CONCAT_SYM
%token<Bytes>  GT_SYM                        /* OPERATOR */
%token<Bytes>  HANDLER_SYM
%token<Bytes>  HASH_SYM
%token<Bytes>  HAVING                        /* SQL-2003-R */
%token<Bytes>  HELP_SYM
%token<Bytes>  HEX_NUM
%token<Bytes>  HIGH_PRIORITY
%token<Bytes>  HOST_SYM
%token<Bytes>  HOSTS_SYM
%token<Bytes>  HOUR_MICROSECOND_SYM
%token<Bytes>  HOUR_MINUTE_SYM
%token<Bytes>  HOUR_SECOND_SYM
%token<Bytes>  HOUR_SYM                      /* SQL-2003-R */
%token<Bytes>  IDENT
%token<Bytes>  IDENTIFIED_SYM
%token<Bytes>  IDENT_QUOTED
%token<Bytes>  IF
%token<Bytes>  IGNORE_SYM
%token<Bytes>  IGNORE_SERVER_IDS_SYM
%token<Bytes>  IMPORT
%token<Bytes>  INDEXES
%token<Bytes>  INDEX_SYM
%token<Bytes>  INFILE
%token<Bytes>  INITIAL_SIZE_SYM
%token<Bytes>  INNER_SYM                     /* SQL-2003-R */
%token<Bytes>  INOUT_SYM                     /* SQL-2003-R */
%token<Bytes>  INSENSITIVE_SYM               /* SQL-2003-R */
%token<Bytes>  INSERT                        /* SQL-2003-R */
%token<Bytes>  INSERT_METHOD
%token<Bytes>  INSTALL_SYM
%token<Bytes>  INTERVAL_SYM                  /* SQL-2003-R */
%token<Bytes>  INTO                          /* SQL-2003-R */
%token<Bytes>  INT_SYM                       /* SQL-2003-R */
%token<Bytes>  INVOKER_SYM
%token<Bytes>  IN_SYM                        /* SQL-2003-R */
%token<Bytes>  IO_AFTER_GTIDS                /* MYSQL, FUTURE-USE */
%token<Bytes>  IO_BEFORE_GTIDS               /* MYSQL, FUTURE-USE */
%token<Bytes>  IO_SYM
%token<Bytes>  IPC_SYM
%token<Bytes>  IS                            /* SQL-2003-R */
%token<Bytes>  ISOLATION                     /* SQL-2003-R */
%token<Bytes>  ISSUER_SYM
%token<Bytes>  ITERATE_SYM
%token<Bytes>  JOIN_SYM                      /* SQL-2003-R */
%token<Bytes>  KEYS
%token<Bytes>  KEY_BLOCK_SIZE
%token<Bytes>  KEY_SYM                       /* SQL-2003-N */
%token<Bytes>  KILL_SYM
%token<Bytes>  LANGUAGE_SYM                  /* SQL-2003-R */
%token<Bytes>  LAST_SYM                      /* SQL-2003-N */
%token<Bytes>  LE                            /* OPERATOR */
%token<Bytes>  LEADING                       /* SQL-2003-R */
%token<Bytes>  LEAVES
%token<Bytes>  LEAVE_SYM
%token<Bytes>  LEFT                          /* SQL-2003-R */
%token<Bytes>  LESS_SYM
%token<Bytes>  LEVEL_SYM
%token<Bytes>  LEX_HOSTNAME
%token<Bytes>  LIKE                          /* SQL-2003-R */
%token<Bytes>  LIMIT
%token<Bytes>  LINEAR_SYM
%token<Bytes>  LINES
%token<Bytes>  LINESTRING
%token<Bytes>  LIST_SYM
%token<Bytes>  LOAD
%token<Bytes>  LOCAL_SYM                     /* SQL-2003-R */
%token<Bytes>  LOCATOR_SYM                   /* SQL-2003-N */
%token<Bytes>  LOCKS_SYM
%token<Bytes>  LOCK_SYM
%token<Bytes>  LOGFILE_SYM
%token<Bytes>  LOGS_SYM
%token<Bytes>  LONGBLOB
%token<Bytes>  LONGTEXT
%token<Bytes>  LONG_NUM
%token<Bytes>  LONG_SYM
%token<Bytes>  LOOP_SYM
%token<Bytes>  LOW_PRIORITY
%token<Bytes>  LT                            /* OPERATOR */
%token<Bytes>  MASTER_AUTO_POSITION_SYM
%token<Bytes>  MASTER_BIND_SYM
%token<Bytes>  MASTER_CONNECT_RETRY_SYM
%token<Bytes>  MASTER_DELAY_SYM
%token<Bytes>  MASTER_HOST_SYM
%token<Bytes>  MASTER_LOG_FILE_SYM
%token<Bytes>  MASTER_LOG_POS_SYM
%token<Bytes>  MASTER_PASSWORD_SYM
%token<Bytes>  MASTER_PORT_SYM
%token<Bytes>  MASTER_RETRY_COUNT_SYM
%token<Bytes>  MASTER_SERVER_ID_SYM
%token<Bytes>  MASTER_SSL_CAPATH_SYM
%token<Bytes>  MASTER_SSL_CA_SYM
%token<Bytes>  MASTER_SSL_CERT_SYM
%token<Bytes>  MASTER_SSL_CIPHER_SYM
%token<Bytes>  MASTER_SSL_CRL_SYM
%token<Bytes>  MASTER_SSL_CRLPATH_SYM
%token<Bytes>  MASTER_SSL_KEY_SYM
%token<Bytes>  MASTER_SSL_SYM
%token<Bytes>  MASTER_SSL_VERIFY_SERVER_CERT_SYM
%token<Bytes>  MASTER_SYM
%token<Bytes>  MASTER_USER_SYM
%token<Bytes>  MASTER_HEARTBEAT_PERIOD_SYM
%token<Bytes>  MATCH                         /* SQL-2003-R */
%token<Bytes>  MAX_CONNECTIONS_PER_HOUR
%token<Bytes>  MAX_QUERIES_PER_HOUR
%token<Bytes>  MAX_ROWS
%token<Bytes>  MAX_SIZE_SYM
%token<Bytes>  MAX_SYM                       /* SQL-2003-N */
%token<Bytes>  MAX_UPDATES_PER_HOUR
%token<Bytes>  MAX_USER_CONNECTIONS_SYM
%token<Bytes>  MAX_VALUE_SYM                 /* SQL-2003-N */
%token<Bytes>  MEDIUMBLOB
%token<Bytes>  MEDIUMINT
%token<Bytes>  MEDIUMTEXT
%token<Bytes>  MEDIUM_SYM
%token<Bytes>  MEMORY_SYM
%token<Bytes>  MERGE_SYM                     /* SQL-2003-R */
%token<Bytes>  MESSAGE_TEXT_SYM              /* SQL-2003-N */
%token<Bytes>  MICROSECOND_SYM               /* MYSQL-FUNC */
%token<Bytes>  MIGRATE_SYM
%token<Bytes>  MINUTE_MICROSECOND_SYM
%token<Bytes>  MINUTE_SECOND_SYM
%token<Bytes>  MINUTE_SYM                    /* SQL-2003-R */
%token<Bytes>  MIN_ROWS
%token<Bytes>  MIN_SYM                       /* SQL-2003-N */
%token<Bytes>  MODE_SYM
%token<Bytes>  MODIFIES_SYM                  /* SQL-2003-R */
%token<Bytes>  MODIFY_SYM
%token<Bytes>  MOD_SYM                       /* SQL-2003-N */
%token<Bytes>  MONTH_SYM                     /* SQL-2003-R */
%token<Bytes>  MULTILINESTRING
%token<Bytes>  MULTIPOINT
%token<Bytes>  MULTIPOLYGON
%token<Bytes>  MUTEX_SYM
%token<Bytes>  MYSQL_ERRNO_SYM
%token<Bytes>  NAMES_SYM                     /* SQL-2003-N */
%token<Bytes>  NAME_SYM                      /* SQL-2003-N */
%token<Bytes>  NATIONAL_SYM                  /* SQL-2003-R */
%token<Bytes>  NATURAL                       /* SQL-2003-R */
%token<Bytes>  NCHAR_STRING
%token<Bytes>  NCHAR_SYM                     /* SQL-2003-R */
%token<Bytes>  NDBCLUSTER_SYM
%token<Bytes>  NE                            /* OPERATOR */
%token<Bytes>  NEG
%token<Bytes>  NEW_SYM                       /* SQL-2003-R */
%token<Bytes>  NEXT_SYM                      /* SQL-2003-N */
%token<Bytes>  NODEGROUP_SYM
%token<Bytes>  NONE_SYM                      /* SQL-2003-R */
%token<Bytes>  NOT2_SYM
%token<Bytes>  NOT_SYM                       /* SQL-2003-R */
%token<Bytes>  NOW_SYM
%token<Bytes>  NO_SYM                        /* SQL-2003-R */
%token<Bytes>  NO_WAIT_SYM
%token<Bytes>  NO_WRITE_TO_BINLOG
%token<Bytes>  NULL_SYM                      /* SQL-2003-R */
%token<Bytes>  NUM
%token<Bytes>  NUMBER_SYM                    /* SQL-2003-N */
%token<Bytes>  NUMERIC_SYM                   /* SQL-2003-R */
%token<Bytes>  NVARCHAR_SYM
%token<Bytes>  OFFSET_SYM
%token<Bytes>  OLD_PASSWORD
%token<Bytes>  ON                            /* SQL-2003-R */
%token<Bytes>  ONE_SYM
%token<Bytes>  ONLY_SYM                      /* SQL-2003-R */
%token<Bytes>  OPEN_SYM                      /* SQL-2003-R */
%token<Bytes>  OPTIMIZE
%token<Bytes>  OPTIONS_SYM
%token<Bytes>  OPTION                        /* SQL-2003-N */
%token<Bytes>  OPTIONALLY
%token<Bytes>  OR2_SYM
%token<Bytes>  ORDER_SYM                     /* SQL-2003-R */
%token<Bytes>  OR_OR_SYM                     /* OPERATOR */
%token<Bytes>  OR_SYM                        /* SQL-2003-R */
%token<Bytes>  OUTER
%token<Bytes>  OUTFILE
%token<Bytes>  OUT_SYM                       /* SQL-2003-R */
%token<Bytes>  OWNER_SYM
%token<Bytes>  PACK_KEYS_SYM
%token<Bytes>  PAGE_SYM
%token<Bytes>  PARAM_MARKER
%token<Bytes>  PARSER_SYM
%token<Bytes>  PARTIAL                       /* SQL-2003-N */
%token<Bytes>  PARTITION_SYM                 /* SQL-2003-R */
%token<Bytes>  PARTITIONS_SYM
%token<Bytes>  PARTITIONING_SYM
%token<Bytes>  PASSWORD
%token<Bytes>  PHASE_SYM
%token<Bytes>  PLUGIN_DIR_SYM                /* INTERNAL */
%token<Bytes>  PLUGIN_SYM
%token<Bytes>  PLUGINS_SYM
%token<Bytes>  POINT_SYM
%token<Bytes>  POLYGON
%token<Bytes>  PORT_SYM
%token<Bytes>  POSITION_SYM                  /* SQL-2003-N */
%token<Bytes>  PRECISION                     /* SQL-2003-R */
%token<Bytes>  PREPARE_SYM                   /* SQL-2003-R */
%token<Bytes>  PRESERVE_SYM
%token<Bytes>  PREV_SYM
%token<Bytes>  PRIMARY_SYM                   /* SQL-2003-R */
%token<Bytes>  PRIVILEGES                    /* SQL-2003-N */
%token<Bytes>  PROCEDURE_SYM                 /* SQL-2003-R */
%token<Bytes>  PROCESS
%token<Bytes>  PROCESSLIST_SYM
%token<Bytes>  PROFILE_SYM
%token<Bytes>  PROFILES_SYM
%token<Bytes>  PROXY_SYM
%token<Bytes>  PURGE
%token<Bytes>  QUARTER_SYM
%token<Bytes>  QUERY_SYM
%token<Bytes>  QUICK
%token<Bytes>  RANGE_SYM                     /* SQL-2003-R */
%token<Bytes>  READS_SYM                     /* SQL-2003-R */
%token<Bytes>  READ_ONLY_SYM
%token<Bytes>  READ_SYM                      /* SQL-2003-N */
%token<Bytes>  READ_WRITE_SYM
%token<Bytes>  REAL                          /* SQL-2003-R */
%token<Bytes>  REBUILD_SYM
%token<Bytes>  RECOVER_SYM
%token<Bytes>  REDOFILE_SYM
%token<Bytes>  REDO_BUFFER_SIZE_SYM
%token<Bytes>  REDUNDANT_SYM
%token<Bytes>  REFERENCES                    /* SQL-2003-R */
%token<Bytes>  REGEXP
%token<Bytes>  RELAY
%token<Bytes>  RELAYLOG_SYM
%token<Bytes>  RELAY_LOG_FILE_SYM
%token<Bytes>  RELAY_LOG_POS_SYM
%token<Bytes>  RELAY_THREAD
%token<Bytes>  RELEASE_SYM                   /* SQL-2003-R */
%token<Bytes>  RELOAD
%token<Bytes>  REMOVE_SYM
%token<Bytes>  RENAME
%token<Bytes>  REORGANIZE_SYM
%token<Bytes>  REPAIR
%token<Bytes>  REPEATABLE_SYM                /* SQL-2003-N */
%token<Bytes>  REPEAT_SYM                    /* MYSQL-FUNC */
%token<Bytes>  REPLACE                       /* MYSQL-FUNC */
%token<Bytes>  REPLICATION
%token<Bytes>  REQUIRE_SYM
%token<Bytes>  RESET_SYM
%token<Bytes>  RESIGNAL_SYM                  /* SQL-2003-R */
%token<Bytes>  RESOURCES
%token<Bytes>  RESTORE_SYM
%token<Bytes>  RESTRICT
%token<Bytes>  RESUME_SYM
%token<Bytes>  RETURNED_SQLSTATE_SYM         /* SQL-2003-N */
%token<Bytes>  RETURNS_SYM                   /* SQL-2003-R */
%token<Bytes>  RETURN_SYM                    /* SQL-2003-R */
%token<Bytes>  REVERSE_SYM
%token<Bytes>  REVOKE                        /* SQL-2003-R */
%token<Bytes>  RIGHT                         /* SQL-2003-R */
%token<Bytes>  ROLLBACK_SYM                  /* SQL-2003-R */
%token<Bytes>  ROLLUP_SYM                    /* SQL-2003-R */
%token<Bytes>  ROUTINE_SYM                   /* SQL-2003-N */
%token<Bytes>  ROWS_SYM                      /* SQL-2003-R */
%token<Bytes>  ROW_FORMAT_SYM
%token<Bytes>  ROW_SYM                       /* SQL-2003-R */
%token<Bytes>  ROW_COUNT_SYM                 /* SQL-2003-N */
%token<Bytes>  RTREE_SYM
%token<Bytes>  SAVEPOINT_SYM                 /* SQL-2003-R */
%token<Bytes>  SCHEDULE_SYM
%token<Bytes>  SCHEMA_NAME_SYM               /* SQL-2003-N */
%token<Bytes>  SECOND_MICROSECOND_SYM
%token<Bytes>  SECOND_SYM                    /* SQL-2003-R */
%token<Bytes>  SECURITY_SYM                  /* SQL-2003-N */
%token<Bytes>  SELECT_SYM                    /* SQL-2003-R */
%token<Bytes>  SENSITIVE_SYM                 /* FUTURE-USE */
%token<Bytes>  SEPARATOR_SYM
%token<Bytes>  SERIALIZABLE_SYM              /* SQL-2003-N */
%token<Bytes>  SERIAL_SYM
%token<Bytes>  SESSION_SYM                   /* SQL-2003-N */
%token<Bytes>  SERVER_SYM
%token<Bytes>  SERVER_OPTIONS
%token<Bytes>  SET                           /* SQL-2003-R */
%token<Bytes>  SET_VAR
%token<Bytes>  SHARE_SYM
%token<Bytes>  SHIFT_LEFT                    /* OPERATOR */
%token<Bytes>  SHIFT_RIGHT                   /* OPERATOR */
%token<Bytes>  SHOW
%token<Bytes>  SHUTDOWN
%token<Bytes>  SIGNAL_SYM                    /* SQL-2003-R */
%token<Bytes>  SIGNED_SYM
%token<Bytes>  SIMPLE_SYM                    /* SQL-2003-N */
%token<Bytes>  SLAVE
%token<Bytes>  SLOW
%token<Bytes>  SMALLINT                      /* SQL-2003-R */
%token<Bytes>  SNAPSHOT_SYM
%token<Bytes>  SOCKET_SYM
%token<Bytes>  SONAME_SYM
%token<Bytes>  SOUNDS_SYM
%token<Bytes>  SOURCE_SYM
%token<Bytes>  SPATIAL_SYM
%token<Bytes>  SPECIFIC_SYM                  /* SQL-2003-R */
%token<Bytes>  SQLEXCEPTION_SYM              /* SQL-2003-R */
%token<Bytes>  SQLSTATE_SYM                  /* SQL-2003-R */
%token<Bytes>  SQLWARNING_SYM                /* SQL-2003-R */
%token<Bytes>  SQL_AFTER_GTIDS               /* MYSQL */
%token<Bytes>  SQL_AFTER_MTS_GAPS            /* MYSQL */
%token<Bytes>  SQL_BEFORE_GTIDS              /* MYSQL */
%token<Bytes>  SQL_BIG_RESULT
%token<Bytes>  SQL_BUFFER_RESULT
%token<Bytes>  SQL_CACHE_SYM
%token<Bytes>  SQL_CALC_FOUND_ROWS
%token<Bytes>  SQL_NO_CACHE_SYM
%token<Bytes>  SQL_SMALL_RESULT
%token<Bytes>  SQL_SYM                       /* SQL-2003-R */
%token<Bytes>  SQL_THREAD
%token<Bytes>  SSL_SYM
%token<Bytes>  STARTING
%token<Bytes>  STARTS_SYM
%token<Bytes>  START_SYM                     /* SQL-2003-R */
%token<Bytes>  STATS_AUTO_RECALC_SYM
%token<Bytes>  STATS_PERSISTENT_SYM
%token<Bytes>  STATS_SAMPLE_PAGES_SYM
%token<Bytes>  STATUS_SYM
%token<Bytes>  STDDEV_SAMP_SYM               /* SQL-2003-N */
%token<Bytes>  STD_SYM
%token<Bytes>  STOP_SYM
%token<Bytes>  STORAGE_SYM
%token<Bytes>  STRAIGHT_JOIN
%token<Bytes>  STRING_SYM
%token<Bytes>  SUBCLASS_ORIGIN_SYM           /* SQL-2003-N */
%token<Bytes>  SUBDATE_SYM
%token<Bytes>  SUBJECT_SYM
%token<Bytes>  SUBPARTITIONS_SYM
%token<Bytes>  SUBPARTITION_SYM
%token<Bytes>  SUBSTRING                     /* SQL-2003-N */
%token<Bytes>  SUM_SYM                       /* SQL-2003-N */
%token<Bytes>  SUPER_SYM
%token<Bytes>  SUSPEND_SYM
%token<Bytes>  SWAPS_SYM
%token<Bytes>  SWITCHES_SYM
%token<Bytes>  SYSDATE
%token<Bytes>  TABLES
%token<Bytes>  TABLESPACE
%token<Bytes>  TABLE_REF_PRIORITY
%token<Bytes>  TABLE_SYM                     /* SQL-2003-R */
%token<Bytes>  TABLE_CHECKSUM_SYM
%token<Bytes>  TABLE_NAME_SYM                /* SQL-2003-N */
%token<Bytes>  TEMPORARY                     /* SQL-2003-N */
%token<Bytes>  TEMPTABLE_SYM
%token<Bytes>  TERMINATED
%token<Bytes>  TEXT_STRING
%token<Bytes>  TEXT_SYM
%token<Bytes>  THAN_SYM
%token<Bytes>  THEN_SYM                      /* SQL-2003-R */
%token<Bytes>  TIMESTAMP                     /* SQL-2003-R */
%token<Bytes>  TIMESTAMP_ADD
%token<Bytes>  TIMESTAMP_DIFF
%token<Bytes>  TIME_SYM                      /* SQL-2003-R */
%token<Bytes>  TINYBLOB
%token<Bytes>  TINYINT
%token<Bytes>  TINYTEXT
%token<Bytes>  TO_SYM                        /* SQL-2003-R */
%token<Bytes>  TRAILING                      /* SQL-2003-R */
%token<Bytes>  TRANSACTION_SYM
%token<Bytes>  TRIGGERS_SYM
%token<Bytes>  TRIGGER_SYM                   /* SQL-2003-R */
%token<Bytes>  TRIM                          /* SQL-2003-N */
%token<Bytes>  TRUE_SYM                      /* SQL-2003-R */
%token<Bytes>  TRUNCATE_SYM
%token<Bytes>  TYPES_SYM
%token<Bytes>  TYPE_SYM                      /* SQL-2003-N */
%token<Bytes>  UDF_RETURNS_SYM
%token<Bytes>  ULONGLONG_NUM
%token<Bytes>  UNCOMMITTED_SYM               /* SQL-2003-N */
%token<Bytes>  UNDEFINED_SYM
%token<Bytes>  UNDERSCORE_CHARSET
%token<Bytes>  UNDOFILE_SYM
%token<Bytes>  UNDO_BUFFER_SIZE_SYM
%token<Bytes>  UNDO_SYM                      /* FUTURE-USE */
%token<Bytes>  UNICODE_SYM
%token<Bytes>  UNINSTALL_SYM
%token<Bytes>  UNION_SYM                     /* SQL-2003-R */
%token<Bytes>  UNIQUE_SYM
%token<Bytes>  UNKNOWN_SYM                   /* SQL-2003-R */
%token<Bytes>  UNLOCK_SYM
%token<Bytes>  UNSIGNED
%token<Bytes>  UNTIL_SYM
%token<Bytes>  UPDATE_SYM                    /* SQL-2003-R */
%token<Bytes>  UPGRADE_SYM
%token<Bytes>  USAGE                         /* SQL-2003-N */
%token<Bytes>  USER                          /* SQL-2003-R */
%token<Bytes>  USE_FRM
%token<Bytes>  USE_SYM
%token<Bytes>  USING                         /* SQL-2003-R */
%token<Bytes>  UTC_DATE_SYM
%token<Bytes>  UTC_TIMESTAMP_SYM
%token<Bytes>  UTC_TIME_SYM
%token<Bytes>  VALUES                        /* SQL-2003-R */
%token<Bytes>  VALUE_SYM                     /* SQL-2003-R */
%token<Bytes>  VARBINARY
%token<Bytes>  VARCHAR                       /* SQL-2003-R */
%token<Bytes>  VARIABLES
%token<Bytes>  VARIANCE_SYM
%token<Bytes>  VARYING                       /* SQL-2003-R */
%token<Bytes>  VAR_SAMP_SYM
%token<Bytes>  VIEW_SYM                      /* SQL-2003-N */
%token<Bytes>  WAIT_SYM
%token<Bytes>  WARNINGS
%token<Bytes>  WEEK_SYM
%token<Bytes>  WEIGHT_STRING_SYM
%token<Bytes>  WHEN_SYM                      /* SQL-2003-R */
%token<Bytes>  WHERE                         /* SQL-2003-R */
%token<Bytes>  WHILE_SYM
%token<Bytes>  WITH                          /* SQL-2003-R */
%token<Bytes>  WITH_CUBE_SYM                 /* INTERNAL */
%token<Bytes>  WITH_ROLLUP_SYM               /* INTERNAL */
%token<Bytes>  WORK_SYM                      /* SQL-2003-N */
%token<Bytes>  WRAPPER_SYM
%token<Bytes>  WRITE_SYM                     /* SQL-2003-N */
%token<Bytes>  X509_SYM
%token<Bytes>  XA_SYM
%token<Bytes>  XML_SYM
%token<Bytes>  XOR
%token<Bytes>  YEAR_MONTH_SYM
%token<Bytes>  YEAR_SYM                      /* SQL-2003-R */
%token<Bytes>  ZEROFILL

%left   JOIN_SYM INNER_SYM STRAIGHT_JOIN CROSS LEFT RIGHT
/* A dummy token to force the priority of table_ref production in a join. */
%left   TABLE_REF_PRIORITY
%left   SET_VAR
%left   OR_OR_SYM OR_SYM OR2_SYM
%left   XOR
%left   AND_SYM AND_AND_SYM
%left   BETWEEN_SYM CASE_SYM WHEN_SYM THEN_SYM ELSE
%left   EQ EQUAL_SYM GE GT_SYM LE LT NE IS LIKE REGEXP IN_SYM
%left   '|'
%left   '&'
%left   SHIFT_LEFT SHIFT_RIGHT
%left   '-' '+'
%left   '*' '/' '%' DIV_SYM MOD_SYM
%left   '^'
%left   NEG '~'
%right  NOT_SYM NOT2_SYM
%right  BINARY COLLATE_SYM
%left  INTERVAL_SYM

%start query
%type <statement> verb_clause statement


/* DDL */
%type <statement> alter create drop rename truncate

/* DML */
%type <statement> select insert update delete replace call do handler load

/* Transaction */
%type <statement> commit lock release rollback savepoint start unlock xa 

/* DAL */
%type <statement> analyze binlog_base64_event check checksum optimize repair flush grant install uninstall kill keycache partition_entry preload reset revoke set show 

/* Replication Statement */
%type <statement> change purge slave 

/* Prepare */
%type <statement> deallocate execute prepare

/* Compound-Statement */
%type <statement> get_diagnostics resignal_stmt signal_stmt

/* MySQL Utility Statement */
%type <statement> describe help use

%type <Bytes> ident IDENT_sys keyword keyword_sp
%type <Table> table_name_with_opt_use_partition table_ident into_table insert_table table_ident_nodb sp_name

%type <empty> '.'

%%


query:
  END_OF_INPUT { SetParseTree(MySQLlex, nil) } 
| verb_clause ';' opt_end_of_input { SetParseTree(MySQLlex, $1) } 
| verb_clause END_OF_INPUT { SetParseTree(MySQLlex, $1) }
; 

opt_end_of_input:
 
| END_OF_INPUT;

verb_clause:
  statement { $$ = $1}
| begin { $$ = &Begin{} };

statement:
  alter {$$ = $1}
| analyze {$$ = $1} 
| binlog_base64_event {$$ = $1}
| call {$$ = $1} 
| change {$$ = $1}
| check {$$ = $1}
| checksum {$$ = $1}
| commit {$$ = $1}
| create {$$ = $1}
| deallocate {$$ = $1}
| delete {$$ = $1}
| describe {$$ = $1}
| do {$$ = $1}
| drop {$$ = $1}
| execute {$$ = $1}
| flush {$$ = $1}
| get_diagnostics {$$ = $1}
| grant {$$ = $1}
| handler {$$ = $1}
| help {$$ = $1}
| insert {$$ = $1}
| install {$$ = $1}
| kill {$$ = $1}
| load {$$ = $1}
| lock {$$ = $1}
| optimize {$$ = $1}
| keycache {$$ = $1}
| partition_entry {$$ = $1}
| preload {$$ = $1}
| prepare {$$ = $1}
| purge {$$ = $1}
| release {$$ = $1}
| rename {$$ = $1}
| repair {$$ = $1}
| replace {$$ = $1}
| reset {$$ = $1}
| resignal_stmt {$$ = $1}
| revoke {$$ = $1}
| rollback {$$ = $1}
| savepoint {$$ = $1}
| select {$$ = $1}
| set {$$ = $1}
| signal_stmt {$$ = $1}
| show {$$ = $1}
| slave {$$ = $1}
| start {$$ = $1}
| truncate {$$ = $1}
| uninstall {$$ = $1}
| unlock {$$ = $1}
| update {$$ = $1}
| use {$$ = $1}
| xa {$$ = $1}
;

deallocate:
  deallocate_or_drop PREPARE_SYM ident { $$ = &Deallocate{} };

deallocate_or_drop:
  DEALLOCATE_SYM
| DROP;

prepare:
  PREPARE_SYM ident FROM prepare_src { $$ = &Prepare{} };

prepare_src:
  TEXT_STRING_sys
| '@' ident_or_text;

execute:
  EXECUTE_SYM ident execute_using { $$ = &Execute{} };

execute_using:
 
| USING execute_var_list;

execute_var_list:
  execute_var_list ',' execute_var_ident
| execute_var_ident;

execute_var_ident:
  '@' ident_or_text;

help:
  HELP_SYM ident_or_text { $$ = &Help{} };

change:
  CHANGE MASTER_SYM TO_SYM master_defs { $$ = &Change{} };

master_defs:
  master_def
| master_defs ',' master_def;

master_def:
  MASTER_HOST_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_BIND_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_USER_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_PASSWORD_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_PORT_SYM EQ ulong_num
| MASTER_CONNECT_RETRY_SYM EQ ulong_num
| MASTER_RETRY_COUNT_SYM EQ ulong_num
| MASTER_DELAY_SYM EQ ulong_num
| MASTER_SSL_SYM EQ ulong_num
| MASTER_SSL_CA_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_SSL_CAPATH_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_SSL_CERT_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_SSL_CIPHER_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_SSL_KEY_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_SSL_VERIFY_SERVER_CERT_SYM EQ ulong_num
| MASTER_SSL_CRL_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_SSL_CRLPATH_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_HEARTBEAT_PERIOD_SYM EQ NUM_literal
| IGNORE_SERVER_IDS_SYM EQ '(' ignore_server_id_list ')'
| MASTER_AUTO_POSITION_SYM EQ ulong_num
| master_file_def;

ignore_server_id_list:
 
| ignore_server_id
| ignore_server_id_list ',' ignore_server_id;

ignore_server_id:
  ulong_num;

master_file_def:
  MASTER_LOG_FILE_SYM EQ TEXT_STRING_sys_nonewline
| MASTER_LOG_POS_SYM EQ ulonglong_num
| RELAY_LOG_FILE_SYM EQ TEXT_STRING_sys_nonewline
| RELAY_LOG_POS_SYM EQ ulong_num;

create:
  CREATE opt_table_options TABLE_SYM opt_if_not_exists table_ident create2 { $$ = &CreateTable{} }
| CREATE opt_unique INDEX_SYM ident key_alg ON table_ident '(' key_list ')' normal_key_options opt_index_lock_algorithm { $$ = &CreateIndex{} }
| CREATE fulltext INDEX_SYM ident init_key_options ON table_ident '(' key_list ')' fulltext_key_options opt_index_lock_algorithm { $$ = &CreateIndex{} }
| CREATE spatial INDEX_SYM ident init_key_options ON table_ident '(' key_list ')' spatial_key_options opt_index_lock_algorithm { $$ = &CreateIndex{} }
| CREATE DATABASE opt_if_not_exists ident opt_create_database_options { $$ = &CreateDatabase{} }
| CREATE view_or_trigger_or_sp_or_event { $$ = &CreateView{} }
| CREATE USER clear_privileges grant_list { $$ = &CreateUser{} }
| CREATE LOGFILE_SYM GROUP_SYM logfile_group_info { $$ = &CreateLog{} }
| CREATE TABLESPACE tablespace_info { $$ = &CreateTablespace{} }
| CREATE server_def { $$ = &CreateServer{} }
;

server_def:
  SERVER_SYM ident_or_text FOREIGN DATA_SYM WRAPPER_SYM ident_or_text OPTIONS_SYM '(' server_options_list ')';

server_options_list:
  server_option
| server_options_list ',' server_option;

server_option:
  USER TEXT_STRING_sys
| HOST_SYM TEXT_STRING_sys
| DATABASE TEXT_STRING_sys
| OWNER_SYM TEXT_STRING_sys
| PASSWORD TEXT_STRING_sys
| SOCKET_SYM TEXT_STRING_sys
| PORT_SYM ulong_num;

event_tail:
  remember_name EVENT_SYM opt_if_not_exists sp_name ON SCHEDULE_SYM ev_schedule_time opt_ev_on_completion opt_ev_status opt_ev_comment DO_SYM ev_sql_stmt;

ev_schedule_time:
  EVERY_SYM expr interval ev_starts ev_ends
| AT_SYM expr;

opt_ev_status:
 
| ENABLE_SYM
| DISABLE_SYM ON SLAVE
| DISABLE_SYM;

ev_starts:
 
| STARTS_SYM expr;

ev_ends:
 
| ENDS_SYM expr;

opt_ev_on_completion:
 
| ev_on_completion;

ev_on_completion:
  ON COMPLETION_SYM PRESERVE_SYM
| ON COMPLETION_SYM NOT_SYM PRESERVE_SYM;

opt_ev_comment:
 
| COMMENT_SYM TEXT_STRING_sys;

ev_sql_stmt:
  ev_sql_stmt_inner;

ev_sql_stmt_inner:
  sp_proc_stmt_statement
| sp_proc_stmt_return
| sp_proc_stmt_if
| case_stmt_specification
| sp_labeled_block
| sp_unlabeled_block
| sp_labeled_control
| sp_proc_stmt_unlabeled
| sp_proc_stmt_leave
| sp_proc_stmt_iterate
| sp_proc_stmt_open
| sp_proc_stmt_fetch
| sp_proc_stmt_close;

clear_privileges:
 ;

sp_name:
  ident '.' ident { $$ = &TableInfo{Qualifier: $1, Name: $3} }
| ident { $$ = &TableInfo{Name: $1} };

sp_a_chistics:
 
| sp_a_chistics sp_chistic;

sp_c_chistics:
 
| sp_c_chistics sp_c_chistic;

sp_chistic:
  COMMENT_SYM TEXT_STRING_sys
| LANGUAGE_SYM SQL_SYM
| NO_SYM SQL_SYM
| CONTAINS_SYM SQL_SYM
| READS_SYM SQL_SYM DATA_SYM
| MODIFIES_SYM SQL_SYM DATA_SYM
| sp_suid;

sp_c_chistic:
  sp_chistic
| DETERMINISTIC_SYM
| not DETERMINISTIC_SYM;

sp_suid:
  SQL_SYM SECURITY_SYM DEFINER_SYM
| SQL_SYM SECURITY_SYM INVOKER_SYM;

call:
  CALL_SYM sp_name opt_sp_cparam_list { $$ = &Call{SpName:$2} };

opt_sp_cparam_list:
 
| '(' opt_sp_cparams ')';

opt_sp_cparams:
 
| sp_cparams;

sp_cparams:
  sp_cparams ',' expr
| expr;

sp_fdparam_list:
 
| sp_fdparams;

sp_fdparams:
  sp_fdparams ',' sp_fdparam
| sp_fdparam;

sp_init_param:
 ;

sp_fdparam:
  ident sp_init_param type_with_opt_collate;

sp_pdparam_list:
 
| sp_pdparams;

sp_pdparams:
  sp_pdparams ',' sp_pdparam
| sp_pdparam;

sp_pdparam:
  sp_opt_inout sp_init_param ident type_with_opt_collate;

sp_opt_inout:
 
| IN_SYM
| OUT_SYM
| INOUT_SYM;

sp_proc_stmts:
 
| sp_proc_stmts sp_proc_stmt ';';

sp_proc_stmts1:
  sp_proc_stmt ';'
| sp_proc_stmts1 sp_proc_stmt ';';

sp_decls:
 
| sp_decls sp_decl ';';

sp_decl:
  DECLARE_SYM sp_decl_idents type_with_opt_collate sp_opt_default
| DECLARE_SYM ident CONDITION_SYM FOR_SYM sp_cond
| DECLARE_SYM sp_handler_type HANDLER_SYM FOR_SYM sp_hcond_list sp_proc_stmt
| DECLARE_SYM ident CURSOR_SYM FOR_SYM select;

sp_handler_type:
  EXIT_SYM
| CONTINUE_SYM;

sp_hcond_list:
  sp_hcond_element
| sp_hcond_list ',' sp_hcond_element;

sp_hcond_element:
  sp_hcond;

sp_cond:
  ulong_num
| sqlstate;

sqlstate:
  SQLSTATE_SYM opt_value TEXT_STRING_literal;

opt_value:
 
| VALUE_SYM;

sp_hcond:
  sp_cond
| ident
| SQLWARNING_SYM
| not FOUND_SYM
| SQLEXCEPTION_SYM;

signal_stmt:
  SIGNAL_SYM signal_value opt_set_signal_information { $$ = &Signal{} } ;

signal_value:
  ident
| sqlstate;

opt_signal_value:
 
| signal_value;

opt_set_signal_information:
 
| SET signal_information_item_list;

signal_information_item_list:
  signal_condition_information_item_name EQ signal_allowed_expr
| signal_information_item_list ',' signal_condition_information_item_name EQ signal_allowed_expr;

signal_allowed_expr:
  literal
| variable
| simple_ident;

signal_condition_information_item_name:
  CLASS_ORIGIN_SYM
| SUBCLASS_ORIGIN_SYM
| CONSTRAINT_CATALOG_SYM
| CONSTRAINT_SCHEMA_SYM
| CONSTRAINT_NAME_SYM
| CATALOG_NAME_SYM
| SCHEMA_NAME_SYM
| TABLE_NAME_SYM
| COLUMN_NAME_SYM
| CURSOR_NAME_SYM
| MESSAGE_TEXT_SYM
| MYSQL_ERRNO_SYM;

resignal_stmt:
  RESIGNAL_SYM opt_signal_value opt_set_signal_information { $$ = &Resignal{} }; 

get_diagnostics:
  GET_SYM which_area DIAGNOSTICS_SYM diagnostics_information { $$ = &Diagnostics{} } ;

which_area:
 
| CURRENT_SYM;

diagnostics_information:
  statement_information
| CONDITION_SYM condition_number condition_information;

statement_information:
  statement_information_item
| statement_information ',' statement_information_item;

statement_information_item:
  simple_target_specification EQ statement_information_item_name;

simple_target_specification:
  ident
| '@' ident_or_text;

statement_information_item_name:
  NUMBER_SYM
| ROW_COUNT_SYM;

condition_number:
  signal_allowed_expr;

condition_information:
  condition_information_item
| condition_information ',' condition_information_item;

condition_information_item:
  simple_target_specification EQ condition_information_item_name;

condition_information_item_name:
  CLASS_ORIGIN_SYM
| SUBCLASS_ORIGIN_SYM
| CONSTRAINT_CATALOG_SYM
| CONSTRAINT_SCHEMA_SYM
| CONSTRAINT_NAME_SYM
| CATALOG_NAME_SYM
| SCHEMA_NAME_SYM
| TABLE_NAME_SYM
| COLUMN_NAME_SYM
| CURSOR_NAME_SYM
| MESSAGE_TEXT_SYM
| MYSQL_ERRNO_SYM
| RETURNED_SQLSTATE_SYM;

sp_decl_idents:
  ident
| sp_decl_idents ',' ident;

sp_opt_default:
 
| DEFAULT expr;

sp_proc_stmt:
  sp_proc_stmt_statement
| sp_proc_stmt_return
| sp_proc_stmt_if
| case_stmt_specification
| sp_labeled_block
| sp_unlabeled_block
| sp_labeled_control
| sp_proc_stmt_unlabeled
| sp_proc_stmt_leave
| sp_proc_stmt_iterate
| sp_proc_stmt_open
| sp_proc_stmt_fetch
| sp_proc_stmt_close;

sp_proc_stmt_if:
  IF sp_if END IF;

sp_proc_stmt_statement:
  statement;

sp_proc_stmt_return:
  RETURN_SYM expr;

sp_proc_stmt_unlabeled:
  sp_unlabeled_control;

sp_proc_stmt_leave:
  LEAVE_SYM label_ident;

sp_proc_stmt_iterate:
  ITERATE_SYM label_ident;

sp_proc_stmt_open:
  OPEN_SYM ident;

sp_proc_stmt_fetch:
  FETCH_SYM sp_opt_fetch_noise ident INTO sp_fetch_list;

sp_proc_stmt_close:
  CLOSE_SYM ident;

sp_opt_fetch_noise:
 
| NEXT_SYM FROM
| FROM;

sp_fetch_list:
  ident
| sp_fetch_list ',' ident;

sp_if:
  expr THEN_SYM sp_proc_stmts1 sp_elseifs;

sp_elseifs:
 
| ELSEIF_SYM sp_if
| ELSE sp_proc_stmts1;

case_stmt_specification:
  simple_case_stmt
| searched_case_stmt;

simple_case_stmt:
  CASE_SYM expr simple_when_clause_list else_clause_opt END CASE_SYM;

searched_case_stmt:
  CASE_SYM searched_when_clause_list else_clause_opt END CASE_SYM;

simple_when_clause_list:
  simple_when_clause
| simple_when_clause_list simple_when_clause;

searched_when_clause_list:
  searched_when_clause
| searched_when_clause_list searched_when_clause;

simple_when_clause:
  WHEN_SYM expr THEN_SYM sp_proc_stmts1;

searched_when_clause:
  WHEN_SYM expr THEN_SYM sp_proc_stmts1;

else_clause_opt:
 
| ELSE sp_proc_stmts1;

sp_labeled_control:
  label_ident ':' sp_unlabeled_control sp_opt_label;

sp_opt_label:
 
| label_ident;

sp_labeled_block:
  label_ident ':' sp_block_content sp_opt_label;

sp_unlabeled_block:
  sp_block_content;

sp_block_content:
  BEGIN_SYM sp_decls sp_proc_stmts END;

sp_unlabeled_control:
  LOOP_SYM sp_proc_stmts1 END LOOP_SYM
| WHILE_SYM expr DO_SYM sp_proc_stmts1 END WHILE_SYM
| REPEAT_SYM sp_proc_stmts1 UNTIL_SYM expr END REPEAT_SYM;

trg_action_time:
  BEFORE_SYM
| AFTER_SYM;

trg_event:
  INSERT
| UPDATE_SYM
| DELETE_SYM;

change_tablespace_access:
  tablespace_name ts_access_mode;

change_tablespace_info:
  tablespace_name CHANGE ts_datafile change_ts_option_list;

tablespace_info:
  tablespace_name ADD ts_datafile opt_logfile_group_name tablespace_option_list;

opt_logfile_group_name:
 
| USE_SYM LOGFILE_SYM GROUP_SYM ident;

alter_tablespace_info:
  tablespace_name ADD ts_datafile alter_tablespace_option_list
| tablespace_name DROP ts_datafile alter_tablespace_option_list;

logfile_group_info:
  logfile_group_name add_log_file logfile_group_option_list;

alter_logfile_group_info:
  logfile_group_name add_log_file alter_logfile_group_option_list;

add_log_file:
  ADD lg_undofile
| ADD lg_redofile;

change_ts_option_list:
  change_ts_options;

change_ts_options:
  change_ts_option
| change_ts_options change_ts_option
| change_ts_options ',' change_ts_option;

change_ts_option:
  opt_ts_initial_size
| opt_ts_autoextend_size
| opt_ts_max_size;

tablespace_option_list:
 
| tablespace_options;

tablespace_options:
  tablespace_option
| tablespace_options tablespace_option
| tablespace_options ',' tablespace_option;

tablespace_option:
  opt_ts_initial_size
| opt_ts_autoextend_size
| opt_ts_max_size
| opt_ts_extent_size
| opt_ts_nodegroup
| opt_ts_engine
| ts_wait
| opt_ts_comment;

alter_tablespace_option_list:
 
| alter_tablespace_options;

alter_tablespace_options:
  alter_tablespace_option
| alter_tablespace_options alter_tablespace_option
| alter_tablespace_options ',' alter_tablespace_option;

alter_tablespace_option:
  opt_ts_initial_size
| opt_ts_autoextend_size
| opt_ts_max_size
| opt_ts_engine
| ts_wait;

logfile_group_option_list:
 
| logfile_group_options;

logfile_group_options:
  logfile_group_option
| logfile_group_options logfile_group_option
| logfile_group_options ',' logfile_group_option;

logfile_group_option:
  opt_ts_initial_size
| opt_ts_undo_buffer_size
| opt_ts_redo_buffer_size
| opt_ts_nodegroup
| opt_ts_engine
| ts_wait
| opt_ts_comment;

alter_logfile_group_option_list:
 
| alter_logfile_group_options;

alter_logfile_group_options:
  alter_logfile_group_option
| alter_logfile_group_options alter_logfile_group_option
| alter_logfile_group_options ',' alter_logfile_group_option;

alter_logfile_group_option:
  opt_ts_initial_size
| opt_ts_engine
| ts_wait;

ts_datafile:
  DATAFILE_SYM TEXT_STRING_sys;

lg_undofile:
  UNDOFILE_SYM TEXT_STRING_sys;

lg_redofile:
  REDOFILE_SYM TEXT_STRING_sys;

tablespace_name:
  ident;

logfile_group_name:
  ident;

ts_access_mode:
  READ_ONLY_SYM
| READ_WRITE_SYM
| NOT_SYM ACCESSIBLE_SYM;

opt_ts_initial_size:
  INITIAL_SIZE_SYM opt_equal size_number;

opt_ts_autoextend_size:
  AUTOEXTEND_SIZE_SYM opt_equal size_number;

opt_ts_max_size:
  MAX_SIZE_SYM opt_equal size_number;

opt_ts_extent_size:
  EXTENT_SIZE_SYM opt_equal size_number;

opt_ts_undo_buffer_size:
  UNDO_BUFFER_SIZE_SYM opt_equal size_number;

opt_ts_redo_buffer_size:
  REDO_BUFFER_SIZE_SYM opt_equal size_number;

opt_ts_nodegroup:
  NODEGROUP_SYM opt_equal real_ulong_num;

opt_ts_comment:
  COMMENT_SYM opt_equal TEXT_STRING_sys;

opt_ts_engine:
  opt_storage ENGINE_SYM opt_equal storage_engines;

ts_wait:
  WAIT_SYM
| NO_WAIT_SYM;

size_number:
  real_ulonglong_num
| IDENT_sys;

create2:
  '(' create2a
| opt_create_table_options opt_create_partitioning create3
| LIKE table_ident
| '(' LIKE table_ident ')';

create2a:
  create_field_list ')' opt_create_table_options opt_create_partitioning create3
| opt_create_partitioning create_select ')' union_opt;

create3:
 
| opt_duplicate opt_as create_select union_clause
| opt_duplicate opt_as '(' create_select ')' union_opt;

opt_create_partitioning:
  opt_partitioning;

opt_partitioning:
 
| partitioning;

partitioning:
  PARTITION_SYM have_partitioning partition;

have_partitioning:
 ;

partition_entry:
  PARTITION_SYM partition { $$ = &Partition{} };

partition:
  BY part_type_def opt_num_parts opt_sub_part part_defs;

part_type_def:
  opt_linear KEY_SYM opt_key_algo '(' part_field_list ')'
| opt_linear HASH_SYM part_func
| RANGE_SYM part_func
| RANGE_SYM part_column_list
| LIST_SYM part_func
| LIST_SYM part_column_list;

opt_linear:
 
| LINEAR_SYM;

opt_key_algo:
 
| ALGORITHM_SYM EQ real_ulong_num;

part_field_list:
 
| part_field_item_list;

part_field_item_list:
  part_field_item
| part_field_item_list ',' part_field_item;

part_field_item:
  ident;

part_column_list:
  COLUMNS '(' part_field_list ')';

part_func:
  '(' remember_name part_func_expr remember_end ')';

sub_part_func:
  '(' remember_name part_func_expr remember_end ')';

opt_num_parts:
 
| PARTITIONS_SYM real_ulong_num;

opt_sub_part:
 
| SUBPARTITION_SYM BY opt_linear HASH_SYM sub_part_func opt_num_subparts
| SUBPARTITION_SYM BY opt_linear KEY_SYM opt_key_algo '(' sub_part_field_list ')' opt_num_subparts;

sub_part_field_list:
  sub_part_field_item
| sub_part_field_list ',' sub_part_field_item;

sub_part_field_item:
  ident;

part_func_expr:
  bit_expr;

opt_num_subparts:
 
| SUBPARTITIONS_SYM real_ulong_num;

part_defs:
 
| '(' part_def_list ')';

part_def_list:
  part_definition
| part_def_list ',' part_definition;

part_definition:
  PARTITION_SYM part_name opt_part_values opt_part_options opt_sub_partition;

part_name:
  ident;

opt_part_values:
 
| VALUES LESS_SYM THAN_SYM part_func_max
| VALUES IN_SYM part_values_in;

part_func_max:
  MAX_VALUE_SYM
| part_value_item;

part_values_in:
  part_value_item
| '(' part_value_list ')';

part_value_list:
  part_value_item
| part_value_list ',' part_value_item;

part_value_item:
  '(' part_value_item_list ')';

part_value_item_list:
  part_value_expr_item
| part_value_item_list ',' part_value_expr_item;

part_value_expr_item:
  MAX_VALUE_SYM
| bit_expr;

opt_sub_partition:
 
| '(' sub_part_list ')';

sub_part_list:
  sub_part_definition
| sub_part_list ',' sub_part_definition;

sub_part_definition:
  SUBPARTITION_SYM sub_name opt_part_options;

sub_name:
  ident_or_text;

opt_part_options:
 
| opt_part_option_list;

opt_part_option_list:
  opt_part_option_list opt_part_option
| opt_part_option;

opt_part_option:
  TABLESPACE opt_equal ident_or_text
| opt_storage ENGINE_SYM opt_equal storage_engines
| NODEGROUP_SYM opt_equal real_ulong_num
| MAX_ROWS opt_equal real_ulonglong_num
| MIN_ROWS opt_equal real_ulonglong_num
| DATA_SYM DIRECTORY_SYM opt_equal TEXT_STRING_sys
| INDEX_SYM DIRECTORY_SYM opt_equal TEXT_STRING_sys
| COMMENT_SYM opt_equal TEXT_STRING_sys;

create_select:
  SELECT_SYM select_options select_item_list opt_select_from;

opt_as:
 
| AS;

opt_create_database_options:
 
| create_database_options;

create_database_options:
  create_database_option
| create_database_options create_database_option;

create_database_option:
  default_collation
| default_charset;

opt_table_options:
 
| table_options;

table_options:
  table_option
| table_option table_options;

table_option:
  TEMPORARY;

opt_if_not_exists:
 
| IF not EXISTS;

opt_create_table_options:
 
| create_table_options;

create_table_options_space_separated:
  create_table_option
| create_table_option create_table_options_space_separated;

create_table_options:
  create_table_option
| create_table_option create_table_options
| create_table_option ',' create_table_options;

create_table_option:
  ENGINE_SYM opt_equal storage_engines
| MAX_ROWS opt_equal ulonglong_num
| MIN_ROWS opt_equal ulonglong_num
| AVG_ROW_LENGTH opt_equal ulong_num
| PASSWORD opt_equal TEXT_STRING_sys
| COMMENT_SYM opt_equal TEXT_STRING_sys
| AUTO_INC opt_equal ulonglong_num
| PACK_KEYS_SYM opt_equal ulong_num
| PACK_KEYS_SYM opt_equal DEFAULT
| STATS_AUTO_RECALC_SYM opt_equal ulong_num
| STATS_AUTO_RECALC_SYM opt_equal DEFAULT
| STATS_PERSISTENT_SYM opt_equal ulong_num
| STATS_PERSISTENT_SYM opt_equal DEFAULT
| STATS_SAMPLE_PAGES_SYM opt_equal ulong_num
| STATS_SAMPLE_PAGES_SYM opt_equal DEFAULT
| CHECKSUM_SYM opt_equal ulong_num
| TABLE_CHECKSUM_SYM opt_equal ulong_num
| DELAY_KEY_WRITE_SYM opt_equal ulong_num
| ROW_FORMAT_SYM opt_equal row_types
| UNION_SYM opt_equal '(' opt_table_list ')'
| default_charset
| default_collation
| INSERT_METHOD opt_equal merge_insert_types
| DATA_SYM DIRECTORY_SYM opt_equal TEXT_STRING_sys
| INDEX_SYM DIRECTORY_SYM opt_equal TEXT_STRING_sys
| TABLESPACE ident
| STORAGE_SYM DISK_SYM
| STORAGE_SYM MEMORY_SYM
| CONNECTION_SYM opt_equal TEXT_STRING_sys
| KEY_BLOCK_SIZE opt_equal ulong_num;

default_charset:
  opt_default charset opt_equal charset_name_or_default;

default_collation:
  opt_default COLLATE_SYM opt_equal collation_name_or_default;

storage_engines:
  ident_or_text;

known_storage_engines:
  ident_or_text;

row_types:
  DEFAULT
| FIXED_SYM
| DYNAMIC_SYM
| COMPRESSED_SYM
| REDUNDANT_SYM
| COMPACT_SYM;

merge_insert_types:
  NO_SYM
| FIRST_SYM
| LAST_SYM;

opt_select_from:
  opt_limit_clause
| select_from select_lock_type;

udf_type:
  STRING_SYM
| REAL
| DECIMAL_SYM
| INT_SYM;

create_field_list:
  field_list;

field_list:
  field_list_item
| field_list ',' field_list_item;

field_list_item:
  column_def
| key_def;

column_def:
  field_spec opt_check_constraint
| field_spec references;

key_def:
  normal_key_type opt_ident key_alg '(' key_list ')' normal_key_options
| fulltext opt_key_or_index opt_ident init_key_options '(' key_list ')' fulltext_key_options
| spatial opt_key_or_index opt_ident init_key_options '(' key_list ')' spatial_key_options
| opt_constraint constraint_key_type opt_ident key_alg '(' key_list ')' normal_key_options
| opt_constraint FOREIGN KEY_SYM opt_ident '(' key_list ')' references
| opt_constraint check_constraint;

opt_check_constraint:
 
| check_constraint;

check_constraint:
  CHECK_SYM '(' expr ')';

opt_constraint:
 
| constraint;

constraint:
  CONSTRAINT opt_ident;

field_spec:
  field_ident type opt_attribute;

type:
  int_type opt_field_length field_options
| real_type opt_precision field_options
| FLOAT_SYM float_options field_options
| BIT_SYM
| BIT_SYM field_length
| BOOL_SYM
| BOOLEAN_SYM
| char field_length opt_binary
| char opt_binary
| nchar field_length opt_bin_mod
| nchar opt_bin_mod
| BINARY field_length
| BINARY
| varchar field_length opt_binary
| nvarchar field_length opt_bin_mod
| VARBINARY field_length
| YEAR_SYM opt_field_length field_options
| DATE_SYM
| TIME_SYM type_datetime_precision
| TIMESTAMP type_datetime_precision
| DATETIME type_datetime_precision
| TINYBLOB
| BLOB_SYM opt_field_length
| spatial_type
| MEDIUMBLOB
| LONGBLOB
| LONG_SYM VARBINARY
| LONG_SYM varchar opt_binary
| TINYTEXT opt_binary
| TEXT_SYM opt_field_length opt_binary
| MEDIUMTEXT opt_binary
| LONGTEXT opt_binary
| DECIMAL_SYM float_options field_options
| NUMERIC_SYM float_options field_options
| FIXED_SYM float_options field_options
| ENUM '(' string_list ')' opt_binary
| SET '(' string_list ')' opt_binary
| LONG_SYM opt_binary
| SERIAL_SYM;

spatial_type:
  GEOMETRY_SYM
| GEOMETRYCOLLECTION
| POINT_SYM
| MULTIPOINT
| LINESTRING
| MULTILINESTRING
| POLYGON
| MULTIPOLYGON;

char:
  CHAR_SYM;

nchar:
  NCHAR_SYM
| NATIONAL_SYM CHAR_SYM;

varchar:
  char VARYING
| VARCHAR;

nvarchar:
  NATIONAL_SYM VARCHAR
| NVARCHAR_SYM
| NCHAR_SYM VARCHAR
| NATIONAL_SYM CHAR_SYM VARYING
| NCHAR_SYM VARYING;

int_type:
  INT_SYM
| TINYINT
| SMALLINT
| MEDIUMINT
| BIGINT;

real_type:
  REAL
| DOUBLE_SYM
| DOUBLE_SYM PRECISION;

float_options:
 
| field_length
| precision;

precision:
  '(' NUM ',' NUM ')';

type_datetime_precision:
 
| '(' NUM ')';

func_datetime_precision:
 
| '(' ')'
| '(' NUM ')';

field_options:
 
| field_opt_list;

field_opt_list:
  field_opt_list field_option
| field_option;

field_option:
  SIGNED_SYM
| UNSIGNED
| ZEROFILL;

field_length:
  '(' LONG_NUM ')'
| '(' ULONGLONG_NUM ')'
| '(' DECIMAL_NUM ')'
| '(' NUM ')';

opt_field_length:
 
| field_length;

opt_precision:
 
| precision;

opt_attribute:
 
| opt_attribute_list;

opt_attribute_list:
  opt_attribute_list attribute
| attribute;

attribute:
  NULL_SYM
| not NULL_SYM
| DEFAULT now_or_signed_literal
| ON UPDATE_SYM now
| AUTO_INC
| SERIAL_SYM DEFAULT VALUE_SYM
| opt_primary KEY_SYM
| UNIQUE_SYM
| UNIQUE_SYM KEY_SYM
| COMMENT_SYM TEXT_STRING_sys
| COLLATE_SYM collation_name
| COLUMN_FORMAT_SYM DEFAULT
| COLUMN_FORMAT_SYM FIXED_SYM
| COLUMN_FORMAT_SYM DYNAMIC_SYM
| STORAGE_SYM DEFAULT
| STORAGE_SYM DISK_SYM
| STORAGE_SYM MEMORY_SYM;

type_with_opt_collate:
  type opt_collate;

now:
  NOW_SYM func_datetime_precision;

now_or_signed_literal:
  now
| signed_literal;

charset:
  CHAR_SYM SET
| CHARSET;

charset_name:
  ident_or_text
| BINARY;

charset_name_or_default:
  charset_name
| DEFAULT;

opt_load_data_charset:
 
| charset charset_name_or_default;

old_or_new_charset_name:
  ident_or_text
| BINARY;

old_or_new_charset_name_or_default:
  old_or_new_charset_name
| DEFAULT;

collation_name:
  ident_or_text;

opt_collate:
 
| COLLATE_SYM collation_name_or_default;

collation_name_or_default:
  collation_name
| DEFAULT;

opt_default:
 
| DEFAULT;

ascii:
  ASCII_SYM
| BINARY ASCII_SYM
| ASCII_SYM BINARY;

unicode:
  UNICODE_SYM
| UNICODE_SYM BINARY
| BINARY UNICODE_SYM;

opt_binary:
 
| ascii
| unicode
| BYTE_SYM
| charset charset_name opt_bin_mod
| BINARY
| BINARY charset charset_name;

opt_bin_mod:
 
| BINARY;

ws_nweights:
  '(' real_ulong_num ')';

ws_level_flag_desc:
  ASC
| DESC;

ws_level_flag_reverse:
  REVERSE_SYM;

ws_level_flags:
 
| ws_level_flag_desc
| ws_level_flag_desc ws_level_flag_reverse
| ws_level_flag_reverse;

ws_level_number:
  real_ulong_num;

ws_level_list_item:
  ws_level_number ws_level_flags;

ws_level_list:
  ws_level_list_item
| ws_level_list ',' ws_level_list_item;

ws_level_range:
  ws_level_number '-' ws_level_number;

ws_level_list_or_range:
  ws_level_list
| ws_level_range;

opt_ws_levels:
 
| LEVEL_SYM ws_level_list_or_range;

opt_primary:
 
| PRIMARY_SYM;

references:
  REFERENCES table_ident opt_ref_list opt_match_clause opt_on_update_delete;

opt_ref_list:
 
| '(' ref_list ')';

ref_list:
  ref_list ',' ident
| ident;

opt_match_clause:
 
| MATCH FULL
| MATCH PARTIAL
| MATCH SIMPLE_SYM;

opt_on_update_delete:
 
| ON UPDATE_SYM delete_option
| ON DELETE_SYM delete_option
| ON UPDATE_SYM delete_option ON DELETE_SYM delete_option
| ON DELETE_SYM delete_option ON UPDATE_SYM delete_option;

delete_option:
  RESTRICT
| CASCADE
| SET NULL_SYM
| NO_SYM ACTION
| SET DEFAULT;

normal_key_type:
  key_or_index;

constraint_key_type:
  PRIMARY_SYM KEY_SYM
| UNIQUE_SYM opt_key_or_index;

key_or_index:
  KEY_SYM
| INDEX_SYM;

opt_key_or_index:
 
| key_or_index;

keys_or_index:
  KEYS
| INDEX_SYM
| INDEXES;

opt_unique:
 
| UNIQUE_SYM;

fulltext:
  FULLTEXT_SYM;

spatial:
  SPATIAL_SYM;

init_key_options:
 ;

key_alg:
  init_key_options
| init_key_options key_using_alg;

normal_key_options:
 
| normal_key_opts;

fulltext_key_options:
 
| fulltext_key_opts;

spatial_key_options:
 
| spatial_key_opts;

normal_key_opts:
  normal_key_opt
| normal_key_opts normal_key_opt;

spatial_key_opts:
  spatial_key_opt
| spatial_key_opts spatial_key_opt;

fulltext_key_opts:
  fulltext_key_opt
| fulltext_key_opts fulltext_key_opt;

key_using_alg:
  USING btree_or_rtree
| TYPE_SYM btree_or_rtree;

all_key_opt:
  KEY_BLOCK_SIZE opt_equal ulong_num
| COMMENT_SYM TEXT_STRING_sys;

normal_key_opt:
  all_key_opt
| key_using_alg;

spatial_key_opt:
  all_key_opt;

fulltext_key_opt:
  all_key_opt
| WITH PARSER_SYM IDENT_sys;

btree_or_rtree:
  BTREE_SYM
| RTREE_SYM
| HASH_SYM;

key_list:
  key_list ',' key_part order_dir
| key_part order_dir;

key_part:
  ident
| ident '(' NUM ')';

opt_ident:
 
| field_ident;

opt_component:
 
| '.' ident;

string_list:
  text_string
| string_list ',' text_string;

alter:
  ALTER opt_ignore TABLE_SYM table_ident alter_commands { $$ = &AlterTable{} }
| ALTER DATABASE ident_or_empty create_database_options { $$ = &AlterDatabase{} }
| ALTER DATABASE ident UPGRADE_SYM DATA_SYM DIRECTORY_SYM NAME_SYM { $$ = &AlterDatabase{} }
| ALTER PROCEDURE_SYM sp_name sp_a_chistics { $$ = &AlterProcedure{} }
| ALTER FUNCTION_SYM sp_name sp_a_chistics { $$ = &AlterFunction{} }
| ALTER view_algorithm definer_opt view_tail { $$ = &AlterView{} }
| ALTER definer_opt view_tail { $$ = &AlterView{} }
| ALTER definer_opt EVENT_SYM sp_name ev_alter_on_schedule_completion opt_ev_rename_to opt_ev_status opt_ev_comment opt_ev_sql_stmt { $$ = &AlterEvent{} }
| ALTER TABLESPACE alter_tablespace_info { $$ = &AlterTablespace{} }
| ALTER LOGFILE_SYM GROUP_SYM alter_logfile_group_info { $$ = &AlterLogfile{} }
| ALTER TABLESPACE change_tablespace_info { $$ = &AlterTablespace{} }
| ALTER TABLESPACE change_tablespace_access { $$ = &AlterTablespace{} }
| ALTER SERVER_SYM ident_or_text OPTIONS_SYM '(' server_options_list ')' { $$ = &AlterServer{} }
| ALTER USER clear_privileges alter_user_list { $$ = &AlterUser{} }
; 

alter_user_list:
  user PASSWORD EXPIRE_SYM
| alter_user_list ',' user PASSWORD EXPIRE_SYM;

ev_alter_on_schedule_completion:
 
| ON SCHEDULE_SYM ev_schedule_time
| ev_on_completion
| ON SCHEDULE_SYM ev_schedule_time ev_on_completion;

opt_ev_rename_to:
 
| RENAME TO_SYM sp_name;

opt_ev_sql_stmt:
 
| DO_SYM ev_sql_stmt;

ident_or_empty:
 
| ident;

alter_commands:
 
| DISCARD TABLESPACE
| IMPORT TABLESPACE
| alter_list opt_partitioning
| alter_list remove_partitioning
| remove_partitioning
| partitioning
| add_partition_rule
| DROP PARTITION_SYM alt_part_name_list
| REBUILD_SYM PARTITION_SYM opt_no_write_to_binlog all_or_alt_part_name_list
| OPTIMIZE PARTITION_SYM opt_no_write_to_binlog all_or_alt_part_name_list opt_no_write_to_binlog
| ANALYZE_SYM PARTITION_SYM opt_no_write_to_binlog all_or_alt_part_name_list
| CHECK_SYM PARTITION_SYM all_or_alt_part_name_list opt_mi_check_type
| REPAIR PARTITION_SYM opt_no_write_to_binlog all_or_alt_part_name_list opt_mi_repair_type
| COALESCE PARTITION_SYM opt_no_write_to_binlog real_ulong_num
| TRUNCATE_SYM PARTITION_SYM all_or_alt_part_name_list
| reorg_partition_rule
| EXCHANGE_SYM PARTITION_SYM alt_part_name_item WITH TABLE_SYM table_ident have_partitioning;

remove_partitioning:
  REMOVE_SYM PARTITIONING_SYM have_partitioning;

all_or_alt_part_name_list:
  ALL
| alt_part_name_list;

add_partition_rule:
  ADD PARTITION_SYM opt_no_write_to_binlog add_part_extra;

add_part_extra:
 
| '(' part_def_list ')'
| PARTITIONS_SYM real_ulong_num;

reorg_partition_rule:
  REORGANIZE_SYM PARTITION_SYM opt_no_write_to_binlog reorg_parts_rule;

reorg_parts_rule:
 
| alt_part_name_list INTO '(' part_def_list ')';

alt_part_name_list:
  alt_part_name_item
| alt_part_name_list ',' alt_part_name_item;

alt_part_name_item:
  ident;

alter_list:
  alter_list_item
| alter_list ',' alter_list_item;

add_column:
  ADD opt_column;

alter_list_item:
  add_column column_def opt_place
| ADD key_def
| add_column '(' create_field_list ')'
| CHANGE opt_column field_ident field_spec opt_place
| MODIFY_SYM opt_column field_ident type opt_attribute opt_place
| DROP opt_column field_ident opt_restrict
| DROP FOREIGN KEY_SYM field_ident
| DROP PRIMARY_SYM KEY_SYM
| DROP key_or_index field_ident
| DISABLE_SYM KEYS
| ENABLE_SYM KEYS
| ALTER opt_column field_ident SET DEFAULT signed_literal
| ALTER opt_column field_ident DROP DEFAULT
| RENAME opt_to table_ident
| CONVERT_SYM TO_SYM charset charset_name_or_default opt_collate
| create_table_options_space_separated
| FORCE_SYM
| alter_order_clause
| alter_algorithm_option
| alter_lock_option;

opt_index_lock_algorithm:
 
| alter_lock_option
| alter_algorithm_option
| alter_lock_option alter_algorithm_option
| alter_algorithm_option alter_lock_option;

alter_algorithm_option:
  ALGORITHM_SYM opt_equal DEFAULT
| ALGORITHM_SYM opt_equal ident;

alter_lock_option:
  LOCK_SYM opt_equal DEFAULT
| LOCK_SYM opt_equal ident;

opt_column:
 
| COLUMN_SYM;

opt_ignore:
 
| IGNORE_SYM;

opt_restrict:
 
| RESTRICT
| CASCADE;

opt_place:
 
| AFTER_SYM ident
| FIRST_SYM;

opt_to:
 
| TO_SYM
| EQ
| AS;

slave:
  START_SYM SLAVE opt_slave_thread_option_list slave_until slave_connection_opts { $$ = &StartSlave{} }
| STOP_SYM SLAVE opt_slave_thread_option_list { $$ = &StopSlave{} }
;

start:
  START_SYM TRANSACTION_SYM opt_start_transaction_option_list { $$ = &Start{} };

opt_start_transaction_option_list:
 
| start_transaction_option_list;

start_transaction_option_list:
  start_transaction_option
| start_transaction_option_list ',' start_transaction_option;

start_transaction_option:
  WITH CONSISTENT_SYM SNAPSHOT_SYM
| READ_SYM ONLY_SYM
| READ_SYM WRITE_SYM;

slave_connection_opts:
  slave_user_name_opt slave_user_pass_opt slave_plugin_auth_opt slave_plugin_dir_opt;

slave_user_name_opt:
 
| USER EQ TEXT_STRING_sys;

slave_user_pass_opt:
 
| PASSWORD EQ TEXT_STRING_sys;

slave_plugin_auth_opt:
 
| DEFAULT_AUTH_SYM EQ TEXT_STRING_sys;

slave_plugin_dir_opt:
 
| PLUGIN_DIR_SYM EQ TEXT_STRING_sys;

opt_slave_thread_option_list:
 
| slave_thread_option_list;

slave_thread_option_list:
  slave_thread_option
| slave_thread_option_list ',' slave_thread_option;

slave_thread_option:
  SQL_THREAD
| RELAY_THREAD;

slave_until:
 
| UNTIL_SYM slave_until_opts;

slave_until_opts:
  master_file_def
| slave_until_opts ',' master_file_def
| SQL_BEFORE_GTIDS EQ TEXT_STRING_sys
| SQL_AFTER_GTIDS EQ TEXT_STRING_sys
| SQL_AFTER_MTS_GAPS;

checksum:
  CHECKSUM_SYM table_or_tables table_list opt_checksum_type { $$ = &CheckSum{} } ;

opt_checksum_type:
 
| QUICK
| EXTENDED_SYM;

repair:
  REPAIR opt_no_write_to_binlog table_or_tables table_list opt_mi_repair_type { $$ = &Repair{} };

opt_mi_repair_type:
 
| mi_repair_types;

mi_repair_types:
  mi_repair_type
| mi_repair_type mi_repair_types;

mi_repair_type:
  QUICK
| EXTENDED_SYM
| USE_FRM;

analyze:
  ANALYZE_SYM opt_no_write_to_binlog table_or_tables table_list { $$ = &Analyze{} };

binlog_base64_event:
  BINLOG_SYM TEXT_STRING_sys { $$ = &Binlog{} };

check:
  CHECK_SYM table_or_tables table_list opt_mi_check_type { $$ = &Check{} } ;

opt_mi_check_type:
 
| mi_check_types;

mi_check_types:
  mi_check_type
| mi_check_type mi_check_types;

mi_check_type:
  QUICK
| FAST_SYM
| MEDIUM_SYM
| EXTENDED_SYM
| CHANGED
| FOR_SYM UPGRADE_SYM;

optimize:
  OPTIMIZE opt_no_write_to_binlog table_or_tables table_list { $$ = &Optimize{} }; 

opt_no_write_to_binlog:
 
| NO_WRITE_TO_BINLOG
| LOCAL_SYM;

rename:
  RENAME table_or_tables table_to_table_list { $$ = &RenameTable{} }
| RENAME USER clear_privileges rename_list { $$ = &RenameUser{} }
; 

rename_list:
  user TO_SYM user
| rename_list ',' user TO_SYM user;

table_to_table_list:
  table_to_table
| table_to_table_list ',' table_to_table;

table_to_table:
  table_ident TO_SYM table_ident;

keycache:
  CACHE_SYM INDEX_SYM keycache_list_or_parts IN_SYM key_cache_name { $$ = &CacheIndex{} };

keycache_list_or_parts:
  keycache_list
| assign_to_keycache_parts;

keycache_list:
  assign_to_keycache
| keycache_list ',' assign_to_keycache;

assign_to_keycache:
  table_ident cache_keys_spec;

assign_to_keycache_parts:
  table_ident adm_partition cache_keys_spec;

key_cache_name:
  ident
| DEFAULT;

preload:
  LOAD INDEX_SYM INTO CACHE_SYM preload_list_or_parts { $$ = &LoadIndex{} }
;

preload_list_or_parts:
  preload_keys_parts
| preload_list;

preload_list:
  preload_keys
| preload_list ',' preload_keys;

preload_keys:
  table_ident cache_keys_spec opt_ignore_leaves;

preload_keys_parts:
  table_ident adm_partition cache_keys_spec opt_ignore_leaves;

adm_partition:
  PARTITION_SYM have_partitioning '(' all_or_alt_part_name_list ')';

cache_keys_spec:
  cache_key_list_or_empty;

cache_key_list_or_empty:
 
| key_or_index '(' opt_key_usage_list ')';

opt_ignore_leaves:
 
| IGNORE_SYM LEAVES;

select:
  select_init { $$ = &Select{} };

select_init:
  SELECT_SYM select_init2
| '(' select_paren ')' union_opt;

select_paren:
  SELECT_SYM select_part2
| '(' select_paren ')';

select_paren_derived:
  SELECT_SYM select_part2_derived
| '(' select_paren_derived ')';

select_init2:
  select_part2 union_clause;

select_part2:
  select_options select_item_list select_into select_lock_type;

select_into:
  opt_order_clause opt_limit_clause
| into
| select_from
| into select_from
| select_from into;

select_from:
  FROM join_table_list where_clause group_clause having_clause opt_order_clause opt_limit_clause procedure_analyse_clause
| FROM DUAL_SYM where_clause opt_limit_clause;

select_options:
 
| select_option_list;

select_option_list:
  select_option_list select_option
| select_option;

select_option:
  query_expression_option
| SQL_NO_CACHE_SYM
| SQL_CACHE_SYM;

select_lock_type:
 
| FOR_SYM UPDATE_SYM
| LOCK_SYM IN_SYM SHARE_SYM MODE_SYM;

select_item_list:
  select_item_list ',' select_item
| select_item
| '*';

select_item:
  remember_name table_wild remember_end
| remember_name expr remember_end select_alias;

remember_name:
 ;

remember_end:
 ;

select_alias:
 
| AS ident
| AS TEXT_STRING_sys
| ident
| TEXT_STRING_sys;

optional_braces:
 
| '(' ')';

expr:
  expr or expr %prec OR_SYM
| expr XOR expr %prec XOR
| expr and expr %prec AND_SYM
| NOT_SYM expr %prec NOT_SYM
| bool_pri IS TRUE_SYM %prec IS
| bool_pri IS not TRUE_SYM %prec IS
| bool_pri IS FALSE_SYM %prec IS
| bool_pri IS not FALSE_SYM %prec IS
| bool_pri IS UNKNOWN_SYM %prec IS
| bool_pri IS not UNKNOWN_SYM %prec IS
| bool_pri;

bool_pri:
  bool_pri IS NULL_SYM %prec IS
| bool_pri IS not NULL_SYM %prec IS
| bool_pri EQUAL_SYM predicate %prec EQUAL_SYM
| bool_pri comp_op predicate %prec EQ
| bool_pri comp_op all_or_any '(' subselect ')' %prec EQ
| predicate;

predicate:
  bit_expr IN_SYM '(' subselect ')'
| bit_expr not IN_SYM '(' subselect ')'
| bit_expr IN_SYM '(' expr ')'
| bit_expr IN_SYM '(' expr ',' expr_list ')'
| bit_expr not IN_SYM '(' expr ')'
| bit_expr not IN_SYM '(' expr ',' expr_list ')'
| bit_expr BETWEEN_SYM bit_expr AND_SYM predicate
| bit_expr not BETWEEN_SYM bit_expr AND_SYM predicate
| bit_expr SOUNDS_SYM LIKE bit_expr
| bit_expr LIKE simple_expr opt_escape
| bit_expr not LIKE simple_expr opt_escape
| bit_expr REGEXP bit_expr
| bit_expr not REGEXP bit_expr
| bit_expr;

bit_expr:
  bit_expr '|' bit_expr %prec '|'
| bit_expr '&' bit_expr %prec '&'
| bit_expr SHIFT_LEFT bit_expr %prec SHIFT_LEFT
| bit_expr SHIFT_RIGHT bit_expr %prec SHIFT_RIGHT
| bit_expr '+' bit_expr %prec '+'
| bit_expr '-' bit_expr %prec '-'
| bit_expr '+' INTERVAL_SYM expr interval %prec '+'
| bit_expr '-' INTERVAL_SYM expr interval %prec '-'
| bit_expr '*' bit_expr %prec '*'
| bit_expr '/' bit_expr %prec '/'
| bit_expr '%' bit_expr %prec '%'
| bit_expr DIV_SYM bit_expr %prec DIV_SYM
| bit_expr MOD_SYM bit_expr %prec MOD_SYM
| bit_expr '^' bit_expr
| simple_expr;

or:
  OR_SYM
| OR2_SYM;

and:
  AND_SYM
| AND_AND_SYM;

not:
  NOT_SYM
| NOT2_SYM;

not2:
  '!'
| NOT2_SYM;

comp_op:
  EQ
| GE
| GT_SYM
| LE
| LT
| NE;

all_or_any:
  ALL
| ANY_SYM;

simple_expr:
  simple_ident
| function_call_keyword
| function_call_nonkeyword
| function_call_generic
| function_call_conflict
| simple_expr COLLATE_SYM ident_or_text %prec NEG
| literal
| param_marker
| variable
| sum_expr
| simple_expr OR_OR_SYM simple_expr
| '+' simple_expr %prec NEG
| '-' simple_expr %prec NEG
| '~' simple_expr %prec NEG
| not2 simple_expr %prec NEG
| '(' subselect ')'
| '(' expr ')'
| '(' expr ',' expr_list ')'
| ROW_SYM '(' expr ',' expr_list ')'
| EXISTS '(' subselect ')'
| '{' ident expr '}'
| MATCH ident_list_arg AGAINST '(' bit_expr fulltext_options ')'
| BINARY simple_expr %prec NEG
| CAST_SYM '(' expr AS cast_type ')'
| CASE_SYM opt_expr when_list opt_else END
| CONVERT_SYM '(' expr ',' cast_type ')'
| CONVERT_SYM '(' expr USING charset_name ')'
| DEFAULT '(' simple_ident ')'
| VALUES '(' simple_ident_nospvar ')'
| INTERVAL_SYM expr interval '+' expr %prec INTERVAL_SYM;

function_call_keyword:
  CHAR_SYM '(' expr_list ')'
| CHAR_SYM '(' expr_list USING charset_name ')'
| CURRENT_USER optional_braces
| DATE_SYM '(' expr ')'
| DAY_SYM '(' expr ')'
| HOUR_SYM '(' expr ')'
| INSERT '(' expr ',' expr ',' expr ',' expr ')'
| INTERVAL_SYM '(' expr ',' expr ')' %prec INTERVAL_SYM
| INTERVAL_SYM '(' expr ',' expr ',' expr_list ')' %prec INTERVAL_SYM
| LEFT '(' expr ',' expr ')'
| MINUTE_SYM '(' expr ')'
| MONTH_SYM '(' expr ')'
| RIGHT '(' expr ',' expr ')'
| SECOND_SYM '(' expr ')'
| TIME_SYM '(' expr ')'
| TIMESTAMP '(' expr ')'
| TIMESTAMP '(' expr ',' expr ')'
| TRIM '(' expr ')'
| TRIM '(' LEADING expr FROM expr ')'
| TRIM '(' TRAILING expr FROM expr ')'
| TRIM '(' BOTH expr FROM expr ')'
| TRIM '(' LEADING FROM expr ')'
| TRIM '(' TRAILING FROM expr ')'
| TRIM '(' BOTH FROM expr ')'
| TRIM '(' expr FROM expr ')'
| USER '(' ')'
| YEAR_SYM '(' expr ')';

function_call_nonkeyword:
  ADDDATE_SYM '(' expr ',' expr ')'
| ADDDATE_SYM '(' expr ',' INTERVAL_SYM expr interval ')'
| CURDATE optional_braces
| CURTIME func_datetime_precision
| DATE_ADD_INTERVAL '(' expr ',' INTERVAL_SYM expr interval ')' %prec INTERVAL_SYM
| DATE_SUB_INTERVAL '(' expr ',' INTERVAL_SYM expr interval ')' %prec INTERVAL_SYM
| EXTRACT_SYM '(' interval FROM expr ')'
| GET_FORMAT '(' date_time_type ',' expr ')'
| now
| POSITION_SYM '(' bit_expr IN_SYM expr ')'
| SUBDATE_SYM '(' expr ',' expr ')'
| SUBDATE_SYM '(' expr ',' INTERVAL_SYM expr interval ')'
| SUBSTRING '(' expr ',' expr ',' expr ')'
| SUBSTRING '(' expr ',' expr ')'
| SUBSTRING '(' expr FROM expr FOR_SYM expr ')'
| SUBSTRING '(' expr FROM expr ')'
| SYSDATE func_datetime_precision
| TIMESTAMP_ADD '(' interval_time_stamp ',' expr ',' expr ')'
| TIMESTAMP_DIFF '(' interval_time_stamp ',' expr ',' expr ')'
| UTC_DATE_SYM optional_braces
| UTC_TIME_SYM func_datetime_precision
| UTC_TIMESTAMP_SYM func_datetime_precision;

function_call_conflict:
  ASCII_SYM '(' expr ')'
| CHARSET '(' expr ')'
| COALESCE '(' expr_list ')'
| COLLATION_SYM '(' expr ')'
| DATABASE '(' ')'
| IF '(' expr ',' expr ',' expr ')'
| FORMAT_SYM '(' expr ',' expr ')'
| FORMAT_SYM '(' expr ',' expr ',' expr ')'
| MICROSECOND_SYM '(' expr ')'
| MOD_SYM '(' expr ',' expr ')'
| OLD_PASSWORD '(' expr ')'
| PASSWORD '(' expr ')'
| QUARTER_SYM '(' expr ')'
| REPEAT_SYM '(' expr ',' expr ')'
| REPLACE '(' expr ',' expr ',' expr ')'
| REVERSE_SYM '(' expr ')'
| ROW_COUNT_SYM '(' ')'
| TRUNCATE_SYM '(' expr ',' expr ')'
| WEEK_SYM '(' expr ')'
| WEEK_SYM '(' expr ',' expr ')'
| WEIGHT_STRING_SYM '(' expr opt_ws_levels ')'
| WEIGHT_STRING_SYM '(' expr AS CHAR_SYM ws_nweights opt_ws_levels ')'
| WEIGHT_STRING_SYM '(' expr AS BINARY ws_nweights ')'
| WEIGHT_STRING_SYM '(' expr ',' ulong_num ',' ulong_num ',' ulong_num ')'
| geometry_function;

geometry_function:
  CONTAINS_SYM '(' expr ',' expr ')'
| GEOMETRYCOLLECTION '(' expr_list ')'
| LINESTRING '(' expr_list ')'
| MULTILINESTRING '(' expr_list ')'
| MULTIPOINT '(' expr_list ')'
| MULTIPOLYGON '(' expr_list ')'
| POINT_SYM '(' expr ',' expr ')'
| POLYGON '(' expr_list ')';

function_call_generic:
  IDENT_sys '(' opt_udf_expr_list ')'
| ident '.' ident '(' opt_expr_list ')';

fulltext_options:
  opt_natural_language_mode opt_query_expansion
| IN_SYM BOOLEAN_SYM MODE_SYM;

opt_natural_language_mode:
 
| IN_SYM NATURAL LANGUAGE_SYM MODE_SYM;

opt_query_expansion:
 
| WITH QUERY_SYM EXPANSION_SYM;

opt_udf_expr_list:
 
| udf_expr_list;

udf_expr_list:
  udf_expr
| udf_expr_list ',' udf_expr;

udf_expr:
  remember_name expr remember_end select_alias;

sum_expr:
  AVG_SYM '(' in_sum_expr ')'
| AVG_SYM '(' DISTINCT in_sum_expr ')'
| BIT_AND '(' in_sum_expr ')'
| BIT_OR '(' in_sum_expr ')'
| BIT_XOR '(' in_sum_expr ')'
| COUNT_SYM '(' opt_all '*' ')'
| COUNT_SYM '(' in_sum_expr ')'
| COUNT_SYM '(' DISTINCT expr_list ')'
| MIN_SYM '(' in_sum_expr ')'
| MIN_SYM '(' DISTINCT in_sum_expr ')'
| MAX_SYM '(' in_sum_expr ')'
| MAX_SYM '(' DISTINCT in_sum_expr ')'
| STD_SYM '(' in_sum_expr ')'
| VARIANCE_SYM '(' in_sum_expr ')'
| STDDEV_SAMP_SYM '(' in_sum_expr ')'
| VAR_SAMP_SYM '(' in_sum_expr ')'
| SUM_SYM '(' in_sum_expr ')'
| SUM_SYM '(' DISTINCT in_sum_expr ')'
| GROUP_CONCAT_SYM '(' opt_distinct expr_list opt_gorder_clause opt_gconcat_separator ')';

variable:
  '@' variable_aux;

variable_aux:
  ident_or_text SET_VAR expr
| ident_or_text
| '@' opt_var_ident_type ident_or_text opt_component;

opt_distinct:
 
| DISTINCT;

opt_gconcat_separator:
 
| SEPARATOR_SYM text_string;

opt_gorder_clause:
 
| ORDER_SYM BY gorder_list;

gorder_list:
  gorder_list ',' order_ident order_dir
| order_ident order_dir;

in_sum_expr:
  opt_all expr;

cast_type:
  BINARY opt_field_length
| CHAR_SYM opt_field_length opt_binary
| NCHAR_SYM opt_field_length
| SIGNED_SYM
| SIGNED_SYM INT_SYM
| UNSIGNED
| UNSIGNED INT_SYM
| DATE_SYM
| TIME_SYM type_datetime_precision
| DATETIME type_datetime_precision
| DECIMAL_SYM float_options;

opt_expr_list:
 
| expr_list;

expr_list:
  expr
| expr_list ',' expr;

ident_list_arg:
  ident_list
| '(' ident_list ')';

ident_list:
  simple_ident
| ident_list ',' simple_ident;

opt_expr:
 
| expr;

opt_else:
 
| ELSE expr;

when_list:
  WHEN_SYM expr THEN_SYM expr
| when_list WHEN_SYM expr THEN_SYM expr;

table_ref:
  table_factor
| join_table;

join_table_list:
  derived_table_list;

esc_table_ref:
  table_ref
| '{' ident table_ref '}';

derived_table_list:
  esc_table_ref
| derived_table_list ',' esc_table_ref;

join_table:
  table_ref normal_join table_ref %prec TABLE_REF_PRIORITY
| table_ref STRAIGHT_JOIN table_factor
| table_ref normal_join table_ref ON expr
| table_ref STRAIGHT_JOIN table_factor ON expr
| table_ref normal_join table_ref USING '(' using_list ')'
| table_ref NATURAL JOIN_SYM table_factor
| table_ref LEFT opt_outer JOIN_SYM table_ref ON expr
| table_ref LEFT opt_outer JOIN_SYM table_factor USING '(' using_list ')'
| table_ref NATURAL LEFT opt_outer JOIN_SYM table_factor
| table_ref RIGHT opt_outer JOIN_SYM table_ref ON expr
| table_ref RIGHT opt_outer JOIN_SYM table_factor USING '(' using_list ')'
| table_ref NATURAL RIGHT opt_outer JOIN_SYM table_factor;

normal_join:
  JOIN_SYM
| INNER_SYM JOIN_SYM
| CROSS JOIN_SYM;

opt_use_partition:
 
| use_partition;

use_partition:
  PARTITION_SYM '(' using_list ')' have_partitioning;

table_factor:
  table_ident opt_use_partition opt_table_alias opt_key_definition
| select_derived_init get_select_lex select_derived2
| '(' get_select_lex select_derived_union ')' opt_table_alias;

select_derived_union:
  select_derived opt_union_order_or_limit
| select_derived_union UNION_SYM union_option query_specification opt_union_order_or_limit;

select_init2_derived:
  select_part2_derived;

select_part2_derived:
  opt_query_expression_options select_item_list opt_select_from select_lock_type;

select_derived:
  get_select_lex derived_table_list;

select_derived2:
  select_options select_item_list opt_select_from;

get_select_lex:
 ;

select_derived_init:
  SELECT_SYM;

opt_outer:
 
| OUTER;

index_hint_clause:
 
| FOR_SYM JOIN_SYM
| FOR_SYM ORDER_SYM BY
| FOR_SYM GROUP_SYM BY;

index_hint_type:
  FORCE_SYM
| IGNORE_SYM;

index_hint_definition:
  index_hint_type key_or_index index_hint_clause '(' key_usage_list ')'
| USE_SYM key_or_index index_hint_clause '(' opt_key_usage_list ')';

index_hints_list:
  index_hint_definition
| index_hints_list index_hint_definition;

opt_index_hints_list:
 
| index_hints_list;

opt_key_definition:
  opt_index_hints_list;

opt_key_usage_list:
 
| key_usage_list;

key_usage_element:
  ident
| PRIMARY_SYM;

key_usage_list:
  key_usage_element
| key_usage_list ',' key_usage_element;

using_list:
  ident
| using_list ',' ident;

interval:
  interval_time_stamp
| DAY_HOUR_SYM
| DAY_MICROSECOND_SYM
| DAY_MINUTE_SYM
| DAY_SECOND_SYM
| HOUR_MICROSECOND_SYM
| HOUR_MINUTE_SYM
| HOUR_SECOND_SYM
| MINUTE_MICROSECOND_SYM
| MINUTE_SECOND_SYM
| SECOND_MICROSECOND_SYM
| YEAR_MONTH_SYM;

interval_time_stamp:
  DAY_SYM
| WEEK_SYM
| HOUR_SYM
| MINUTE_SYM
| MONTH_SYM
| QUARTER_SYM
| SECOND_SYM
| MICROSECOND_SYM
| YEAR_SYM;

date_time_type:
  DATE_SYM
| TIME_SYM
| TIMESTAMP
| DATETIME;

table_alias:
 
| AS
| EQ;

opt_table_alias:
 
| table_alias ident;

opt_all:
 
| ALL;

where_clause:
 
| WHERE expr;

having_clause:
 
| HAVING expr;

opt_escape:
  ESCAPE_SYM simple_expr
|;

group_clause:
 
| GROUP_SYM BY group_list olap_opt;

group_list:
  group_list ',' order_ident order_dir
| order_ident order_dir;

olap_opt:
 
| WITH_CUBE_SYM
| WITH_ROLLUP_SYM;

alter_order_clause:
  ORDER_SYM BY alter_order_list;

alter_order_list:
  alter_order_list ',' alter_order_item
| alter_order_item;

alter_order_item:
  simple_ident_nospvar order_dir;

opt_order_clause:
 
| order_clause;

order_clause:
  ORDER_SYM BY order_list;

order_list:
  order_list ',' order_ident order_dir
| order_ident order_dir;

order_dir:
 
| ASC
| DESC;

opt_limit_clause_init:
 
| limit_clause;

opt_limit_clause:
 
| limit_clause;

limit_clause:
  LIMIT limit_options;

limit_options:
  limit_option
| limit_option ',' limit_option
| limit_option OFFSET_SYM limit_option;

limit_option:
  ident
| param_marker
| ULONGLONG_NUM
| LONG_NUM
| NUM;

delete_limit_clause:
 
| LIMIT limit_option;

ulong_num:
  NUM
| HEX_NUM
| LONG_NUM
| ULONGLONG_NUM
| DECIMAL_NUM
| FLOAT_NUM;

real_ulong_num:
  NUM
| HEX_NUM
| LONG_NUM
| ULONGLONG_NUM
| dec_num_error;

ulonglong_num:
  NUM
| ULONGLONG_NUM
| LONG_NUM
| DECIMAL_NUM
| FLOAT_NUM;

real_ulonglong_num:
  NUM
| ULONGLONG_NUM
| LONG_NUM
| dec_num_error;

dec_num_error:
  dec_num;

dec_num:
  DECIMAL_NUM
| FLOAT_NUM;

procedure_analyse_clause:
 
| PROCEDURE_SYM ANALYSE_SYM '(' opt_procedure_analyse_params ')';

opt_procedure_analyse_params:
 
| procedure_analyse_param
| procedure_analyse_param ',' procedure_analyse_param;

procedure_analyse_param:
  NUM;

select_var_list_init:
  select_var_list;

select_var_list:
  select_var_list ',' select_var_ident
| select_var_ident;

select_var_ident:
  '@' ident_or_text
| ident_or_text;

into:
  INTO into_destination;

into_destination:
  OUTFILE TEXT_STRING_filesystem opt_load_data_charset opt_field_term opt_line_term
| DUMPFILE TEXT_STRING_filesystem
| select_var_list_init;

do:
  DO_SYM expr_list { $$ = &Do{} };

drop:
  DROP opt_temporary table_or_tables if_exists table_list opt_restrict { $$ = &DropTables{} }
| DROP INDEX_SYM ident ON table_ident opt_index_lock_algorithm { $$ = &DropIndex{} }
| DROP DATABASE if_exists ident { $$ = &DropDatabase{} }
| DROP FUNCTION_SYM if_exists ident '.' ident { $$ = &DropFunction{} }
| DROP FUNCTION_SYM if_exists ident { $$ = &DropFunction{} }
| DROP PROCEDURE_SYM if_exists sp_name { $$ = &DropProcedure{} }
| DROP USER clear_privileges user_list { $$ = &DropUser{} }
| DROP VIEW_SYM if_exists table_list opt_restrict { $$ = &DropView{} }
| DROP EVENT_SYM if_exists sp_name { $$ = &DropEvent{} }
| DROP TRIGGER_SYM if_exists sp_name { $$ = &DropTrigger{} }
| DROP TABLESPACE tablespace_name drop_ts_options_list { $$ = &DropTablespace{} }
| DROP LOGFILE_SYM GROUP_SYM logfile_group_name drop_ts_options_list { $$ = &DropLogfile{} }
| DROP SERVER_SYM if_exists ident_or_text { $$ = &DropServer{} }
;

table_list:
  table_name
| table_list ',' table_name;

table_name:
  table_ident;

table_name_with_opt_use_partition:
  table_ident opt_use_partition { $$ = $1 };

table_alias_ref_list:
  table_alias_ref
| table_alias_ref_list ',' table_alias_ref;

table_alias_ref:
  table_ident_opt_wild;

if_exists:
 
| IF EXISTS;

opt_temporary:
 
| TEMPORARY;

drop_ts_options_list:
 
| drop_ts_options;

drop_ts_options:
  drop_ts_option
| drop_ts_options drop_ts_option
| drop_ts_options_list ',' drop_ts_option;

drop_ts_option:
  opt_ts_engine
| ts_wait;

insert:
  INSERT insert_lock_option opt_ignore into_table insert_field_spec opt_insert_update
  {
    $$ = &Insert{Table: $4}
  }
;

replace:
  REPLACE replace_lock_option into_table insert_field_spec { $$ = &Replace{} };

insert_lock_option:
 
| LOW_PRIORITY
| DELAYED_SYM
| HIGH_PRIORITY;

replace_lock_option:
  opt_low_priority
| DELAYED_SYM;

into_table:
  INTO insert_table { $$ = $2 }
| insert_table { $$ = $1 };

insert_table:
  table_name_with_opt_use_partition { $$ = $1 };

insert_field_spec:
  insert_values
| '(' ')' insert_values
| '(' fields ')' insert_values
| SET ident_eq_list;

fields:
  fields ',' insert_ident
| insert_ident;

insert_values:
  VALUES values_list
| VALUE_SYM values_list
| create_select union_clause
| '(' create_select ')' union_opt;

values_list:
  values_list ',' no_braces
| no_braces;

ident_eq_list:
  ident_eq_list ',' ident_eq_value
| ident_eq_value;

ident_eq_value:
  simple_ident_nospvar equal expr_or_default;

equal:
  EQ
| SET_VAR;

opt_equal:
 
| equal;

no_braces:
  '(' opt_values ')';

opt_values:
 
| values;

values:
  values ',' expr_or_default
| expr_or_default;

expr_or_default:
  expr
| DEFAULT;

opt_insert_update:
 
| ON DUPLICATE_SYM KEY_SYM UPDATE_SYM insert_update_list;

update:
  UPDATE_SYM opt_low_priority opt_ignore join_table_list SET update_list where_clause opt_order_clause delete_limit_clause { $$ = &Update{} };

update_list:
  update_list ',' update_elem
| update_elem;

update_elem:
  simple_ident_nospvar equal expr_or_default;

insert_update_list:
  insert_update_list ',' insert_update_elem
| insert_update_elem;

insert_update_elem:
  simple_ident_nospvar equal expr_or_default;

opt_low_priority:
 
| LOW_PRIORITY;

delete:
  DELETE_SYM opt_delete_options single_multi { $$ = &Delete{} }
;

single_multi:
  FROM table_ident opt_use_partition where_clause opt_order_clause delete_limit_clause
| table_wild_list FROM join_table_list where_clause
| FROM table_alias_ref_list USING join_table_list where_clause;

table_wild_list:
  table_wild_one
| table_wild_list ',' table_wild_one;

table_wild_one:
  ident opt_wild
| ident '.' ident opt_wild;

opt_wild:
 
| '.' '*';

opt_delete_options:
 
| opt_delete_option opt_delete_options;

opt_delete_option:
  QUICK
| LOW_PRIORITY
| IGNORE_SYM;

truncate:
  TRUNCATE_SYM opt_table_sym table_name { $$ = &TruncateTable{} }
;

opt_table_sym:
 
| TABLE_SYM;

opt_profile_defs:
 
| profile_defs;

profile_defs:
  profile_def
| profile_defs ',' profile_def;

profile_def:
  CPU_SYM
| MEMORY_SYM
| BLOCK_SYM IO_SYM
| CONTEXT_SYM SWITCHES_SYM
| PAGE_SYM FAULTS_SYM
| IPC_SYM
| SWAPS_SYM
| SOURCE_SYM
| ALL;

opt_profile_args:
 
| FOR_SYM QUERY_SYM NUM;

show:
  SHOW show_param { $$ = &Show{} };

show_param:
  DATABASES wild_and_where
| opt_full TABLES opt_db wild_and_where
| opt_full TRIGGERS_SYM opt_db wild_and_where
| EVENTS_SYM opt_db wild_and_where
| TABLE_SYM STATUS_SYM opt_db wild_and_where
| OPEN_SYM TABLES opt_db wild_and_where
| PLUGINS_SYM
| ENGINE_SYM known_storage_engines show_engine_param
| ENGINE_SYM ALL show_engine_param
| opt_full COLUMNS from_or_in table_ident opt_db wild_and_where
| master_or_binary LOGS_SYM
| SLAVE HOSTS_SYM
| BINLOG_SYM EVENTS_SYM binlog_in binlog_from opt_limit_clause_init
| RELAYLOG_SYM EVENTS_SYM binlog_in binlog_from opt_limit_clause_init
| keys_or_index from_or_in table_ident opt_db where_clause
| opt_storage ENGINES_SYM
| PRIVILEGES
| COUNT_SYM '(' '*' ')' WARNINGS
| COUNT_SYM '(' '*' ')' ERRORS
| WARNINGS opt_limit_clause_init
| ERRORS opt_limit_clause_init
| PROFILES_SYM
| PROFILE_SYM opt_profile_defs opt_profile_args opt_limit_clause_init
| opt_var_type STATUS_SYM wild_and_where
| opt_full PROCESSLIST_SYM
| opt_var_type VARIABLES wild_and_where
| charset wild_and_where
| COLLATION_SYM wild_and_where
| GRANTS
| GRANTS FOR_SYM user
| CREATE DATABASE opt_if_not_exists ident
| CREATE TABLE_SYM table_ident
| CREATE VIEW_SYM table_ident
| MASTER_SYM STATUS_SYM
| SLAVE STATUS_SYM
| CREATE PROCEDURE_SYM sp_name
| CREATE FUNCTION_SYM sp_name
| CREATE TRIGGER_SYM sp_name
| PROCEDURE_SYM STATUS_SYM wild_and_where
| FUNCTION_SYM STATUS_SYM wild_and_where
| PROCEDURE_SYM CODE_SYM sp_name
| FUNCTION_SYM CODE_SYM sp_name
| CREATE EVENT_SYM sp_name;

show_engine_param:
  STATUS_SYM
| MUTEX_SYM
| LOGS_SYM;

master_or_binary:
  MASTER_SYM
| BINARY;

opt_storage:
 
| STORAGE_SYM;

opt_db:
 
| from_or_in ident;

opt_full:
 
| FULL;

from_or_in:
  FROM
| IN_SYM;

binlog_in:
 
| IN_SYM TEXT_STRING_sys;

binlog_from:
 
| FROM ulonglong_num;

wild_and_where:
 
| LIKE TEXT_STRING_sys
| WHERE expr;

describe:
  describe_command table_ident opt_describe_column { $$ = &Describe{} }
| describe_command opt_extended_describe explanable_command { $$ = &Describe{} }
;

explanable_command:
  select
| insert
| replace
| update
| delete;

describe_command:
  DESC
| DESCRIBE;

opt_extended_describe:
 
| EXTENDED_SYM
| PARTITIONS_SYM
| FORMAT_SYM EQ ident_or_text;

opt_describe_column:
 
| text_string
| ident;

flush:
  FLUSH_SYM opt_no_write_to_binlog flush_options { $$ = &Flush{} };

flush_options:
  table_or_tables opt_table_list opt_flush_lock
| flush_options_list;

opt_flush_lock:
 
| WITH READ_SYM LOCK_SYM
| FOR_SYM EXPORT_SYM;

flush_options_list:
  flush_options_list ',' flush_option
| flush_option;

flush_option:
  ERROR_SYM LOGS_SYM
| ENGINE_SYM LOGS_SYM
| GENERAL LOGS_SYM
| SLOW LOGS_SYM
| BINARY LOGS_SYM
| RELAY LOGS_SYM
| QUERY_SYM CACHE_SYM
| HOSTS_SYM
| PRIVILEGES
| LOGS_SYM
| STATUS_SYM
| DES_KEY_FILE
| RESOURCES;

opt_table_list:
 
| table_list;

reset:
  RESET_SYM reset_options { $$ = &Reset{} };

reset_options:
  reset_options ',' reset_option
| reset_option;

reset_option:
  SLAVE slave_reset_options
| MASTER_SYM
| QUERY_SYM CACHE_SYM;

slave_reset_options:
 
| ALL;

purge:
  PURGE purge_options { $$ = &Purge{} };

purge_options:
  master_or_binary LOGS_SYM purge_option;

purge_option:
  TO_SYM TEXT_STRING_sys
| BEFORE_SYM expr;

kill:
  KILL_SYM kill_option expr { $$ = &Kill{} };

kill_option:
 
| CONNECTION_SYM
| QUERY_SYM;

use:
  USE_SYM ident { $$ = &Use{} };

load:
  LOAD data_or_xml load_data_lock opt_local INFILE TEXT_STRING_filesystem opt_duplicate INTO TABLE_SYM table_ident opt_use_partition opt_load_data_charset opt_xml_rows_identified_by opt_field_term opt_line_term opt_ignore_lines opt_field_or_var_spec opt_load_data_set_spec
  {
    $$ = &Load{}
  }  
;

data_or_xml:
  DATA_SYM
| XML_SYM;

opt_local:
 
| LOCAL_SYM;

load_data_lock:
 
| CONCURRENT
| LOW_PRIORITY;

opt_duplicate:
 
| REPLACE
| IGNORE_SYM;

opt_field_term:
 
| COLUMNS field_term_list;

field_term_list:
  field_term_list field_term
| field_term;

field_term:
  TERMINATED BY text_string
| OPTIONALLY ENCLOSED BY text_string
| ENCLOSED BY text_string
| ESCAPED BY text_string;

opt_line_term:
 
| LINES line_term_list;

line_term_list:
  line_term_list line_term
| line_term;

line_term:
  TERMINATED BY text_string
| STARTING BY text_string;

opt_xml_rows_identified_by:
 
| ROWS_SYM IDENTIFIED_SYM BY text_string;

opt_ignore_lines:
 
| IGNORE_SYM NUM lines_or_rows;

lines_or_rows:
  LINES
| ROWS_SYM;

opt_field_or_var_spec:
 
| '(' fields_or_vars ')'
| '(' ')';

fields_or_vars:
  fields_or_vars ',' field_or_var
| field_or_var;

field_or_var:
  simple_ident_nospvar
| '@' ident_or_text;

opt_load_data_set_spec:
 
| SET load_data_set_list;

load_data_set_list:
  load_data_set_list ',' load_data_set_elem
| load_data_set_elem;

load_data_set_elem:
  simple_ident_nospvar equal remember_name expr_or_default remember_end;

text_literal:
  TEXT_STRING
| NCHAR_STRING
| UNDERSCORE_CHARSET TEXT_STRING
| text_literal TEXT_STRING_literal;

text_string:
  TEXT_STRING_literal
| HEX_NUM
| BIN_NUM;

param_marker:
  PARAM_MARKER;

signed_literal:
  literal
| '+' NUM_literal
| '-' NUM_literal;

literal:
  text_literal
| NUM_literal
| temporal_literal
| NULL_SYM
| FALSE_SYM
| TRUE_SYM
| HEX_NUM
| BIN_NUM
| UNDERSCORE_CHARSET HEX_NUM
| UNDERSCORE_CHARSET BIN_NUM;

NUM_literal:
  NUM
| LONG_NUM
| ULONGLONG_NUM
| DECIMAL_NUM
| FLOAT_NUM;

temporal_literal:
  DATE_SYM TEXT_STRING
| TIME_SYM TEXT_STRING
| TIMESTAMP TEXT_STRING;

insert_ident:
  simple_ident_nospvar
| table_wild;

table_wild:
  ident '.' '*'
| ident '.' ident '.' '*';

order_ident:
  expr;

simple_ident:
  ident
| simple_ident_q;

simple_ident_nospvar:
  ident
| simple_ident_q;

simple_ident_q:
  ident '.' ident
| '.' ident '.' ident
| ident '.' ident '.' ident;

field_ident:
  ident
| ident '.' ident '.' ident
| ident '.' ident
| '.' ident;

table_ident:
  ident { $$ = &TableInfo{Name: $1} }
| ident '.' ident { $$ = &TableInfo{Qualifier: $1, Name: $3} }
| '.' ident { $$ = &TableInfo{Name: $2} } ;

table_ident_opt_wild:
  ident opt_wild
| ident '.' ident opt_wild;

table_ident_nodb:
  ident { $$ = &TableInfo{Name: $1} };

IDENT_sys:
  IDENT { $$ = $1 }
| IDENT_QUOTED { $$ = $1 };

TEXT_STRING_sys_nonewline:
  TEXT_STRING_sys;

TEXT_STRING_sys:
  TEXT_STRING;

TEXT_STRING_literal:
  TEXT_STRING;

TEXT_STRING_filesystem:
  TEXT_STRING;

ident:
  IDENT_sys { $$ = $1 }
| keyword { $$ = $1 };

label_ident:
  IDENT_sys
| keyword_sp;

ident_or_text:
  ident
| TEXT_STRING_sys
| LEX_HOSTNAME;

user:
  ident_or_text
| ident_or_text '@' ident_or_text
| CURRENT_USER optional_braces;

keyword:
  keyword_sp { $$ = $1 }
| ASCII_SYM { $$ = $1 }
| BACKUP_SYM { $$ = $1 }
| BEGIN_SYM { $$ = $1 }
| BYTE_SYM { $$ = $1 }
| CACHE_SYM { $$ = $1 }
| CHARSET { $$ = $1 }
| CHECKSUM_SYM { $$ = $1 }
| CLOSE_SYM { $$ = $1 }
| COMMENT_SYM { $$ = $1 }
| COMMIT_SYM { $$ = $1 }
| CONTAINS_SYM { $$ = $1 }
| DEALLOCATE_SYM { $$ = $1 }
| DO_SYM { $$ = $1 }
| END { $$ = $1 }
| EXECUTE_SYM { $$ = $1 }
| FLUSH_SYM { $$ = $1 }
| FORMAT_SYM { $$ = $1 }
| HANDLER_SYM { $$ = $1 }
| HELP_SYM { $$ = $1 }
| HOST_SYM { $$ = $1 }
| INSTALL_SYM { $$ = $1 }
| LANGUAGE_SYM { $$ = $1 }
| NO_SYM { $$ = $1 }
| OPEN_SYM { $$ = $1 }
| OPTIONS_SYM { $$ = $1 }
| OWNER_SYM { $$ = $1 }
| PARSER_SYM { $$ = $1 } 
| PORT_SYM { $$ = $1 }
| PREPARE_SYM { $$ = $1 }
| REMOVE_SYM { $$ = $1 }
| REPAIR { $$ = $1 }
| RESET_SYM { $$ = $1 }
| RESTORE_SYM { $$ = $1 }
| ROLLBACK_SYM { $$ = $1 }
| SAVEPOINT_SYM { $$ = $1 }
| SECURITY_SYM { $$ = $1 }
| SERVER_SYM { $$ = $1 }
| SIGNED_SYM { $$ = $1 }
| SOCKET_SYM { $$ = $1 }
| SLAVE { $$ = $1 }
| SONAME_SYM { $$ = $1 }
| START_SYM { $$ = $1 }
| STOP_SYM { $$ = $1 }
| TRUNCATE_SYM { $$ = $1 }
| UNICODE_SYM { $$ = $1 }
| UNINSTALL_SYM { $$ = $1 }
| WRAPPER_SYM { $$ = $1 }
| XA_SYM { $$ = $1 }
| UPGRADE_SYM { $$ = $1 }
;

keyword_sp:
  ACTION { $$ = $1 }
| ADDDATE_SYM { $$ = $1 }
| AFTER_SYM { $$ = $1 }
| AGAINST { $$ = $1 }
| AGGREGATE_SYM { $$ = $1 }
| ALGORITHM_SYM { $$ = $1 }
| ANALYSE_SYM { $$ = $1 }
| ANY_SYM { $$ = $1 }
| AT_SYM { $$ = $1 } 
| AUTO_INC { $$ = $1 }
| AUTOEXTEND_SIZE_SYM { $$ = $1 }
| AVG_ROW_LENGTH { $$ = $1 }
| AVG_SYM { $$ = $1 }
| BINLOG_SYM { $$ = $1 }
| BIT_SYM { $$ = $1 }
| BLOCK_SYM { $$ = $1 }
| BOOL_SYM { $$ = $1 }
| BOOLEAN_SYM { $$ = $1 }
| BTREE_SYM { $$ = $1 }
| CASCADED { $$ = $1 }
| CATALOG_NAME_SYM { $$ = $1 }
| CHAIN_SYM { $$ = $1 }
| CHANGED { $$ = $1 }
| CIPHER_SYM { $$ = $1 }
| CLIENT_SYM { $$ = $1 }
| CLASS_ORIGIN_SYM { $$ = $1 }
| COALESCE { $$ = $1 }
| CODE_SYM { $$ = $1 }
| COLLATION_SYM { $$ = $1 }
| COLUMN_NAME_SYM { $$ = $1 }
| COLUMN_FORMAT_SYM { $$ = $1 }
| COLUMNS { $$ = $1 }
| COMMITTED_SYM { $$ = $1 }
| COMPACT_SYM { $$ = $1 }
| COMPLETION_SYM { $$ = $1 }
| COMPRESSED_SYM { $$ = $1 }
| CONCURRENT { $$ = $1 }
| CONNECTION_SYM { $$ = $1 }
| CONSISTENT_SYM { $$ = $1 }
| CONSTRAINT_CATALOG_SYM { $$ = $1 }
| CONSTRAINT_SCHEMA_SYM { $$ = $1 }
| CONSTRAINT_NAME_SYM { $$ = $1 }
| CONTEXT_SYM { $$ = $1 }
| CPU_SYM { $$ = $1 }
| CUBE_SYM { $$ = $1 }
| CURRENT_SYM { $$ = $1 }
| CURSOR_NAME_SYM { $$ = $1 }
| DATA_SYM { $$ = $1 }
| DATAFILE_SYM { $$ = $1 }
| DATETIME { $$ = $1 }
| DATE_SYM { $$ = $1 }
| DAY_SYM { $$ = $1 }
| DEFAULT_AUTH_SYM { $$ = $1 }
| DEFINER_SYM { $$ = $1 }
| DELAY_KEY_WRITE_SYM { $$ = $1 }
| DES_KEY_FILE { $$ = $1 }
| DIAGNOSTICS_SYM { $$ = $1 }
| DIRECTORY_SYM { $$ = $1 }
| DISABLE_SYM { $$ = $1 }
| DISCARD { $$ = $1 }
| DISK_SYM { $$ = $1 }
| DUMPFILE { $$ = $1 }
| DUPLICATE_SYM { $$ = $1 }
| DYNAMIC_SYM { $$ = $1 }
| ENDS_SYM { $$ = $1 }
| ENUM { $$ = $1 }
| ENGINE_SYM { $$ = $1 }
| ENGINES_SYM { $$ = $1 }
| ERROR_SYM { $$ = $1 }
| ERRORS { $$ = $1 }
| ESCAPE_SYM { $$ = $1 }
| EVENT_SYM { $$ = $1 }
| EVENTS_SYM { $$ = $1 }
| EVERY_SYM { $$ = $1 }
| EXCHANGE_SYM { $$ = $1 }
| EXPANSION_SYM { $$ = $1 }
| EXPIRE_SYM { $$ = $1 }
| EXPORT_SYM { $$ = $1 }
| EXTENDED_SYM { $$ = $1 }
| EXTENT_SIZE_SYM { $$ = $1 }
| FAULTS_SYM { $$ = $1 }
| FAST_SYM { $$ = $1 }
| FOUND_SYM { $$ = $1 }
| ENABLE_SYM { $$ = $1 }
| FULL { $$ = $1 }
| FILE_SYM { $$ = $1 }
| FIRST_SYM { $$ = $1 }
| FIXED_SYM { $$ = $1 }
| GENERAL { $$ = $1 }
| GEOMETRY_SYM { $$ = $1 }
| GEOMETRYCOLLECTION { $$ = $1 }
| GET_FORMAT { $$ = $1 }
| GRANTS { $$ = $1 }
| GLOBAL_SYM { $$ = $1 }
| HASH_SYM { $$ = $1 }
| HOSTS_SYM { $$ = $1 }
| HOUR_SYM { $$ = $1 }
| IDENTIFIED_SYM { $$ = $1 }
| IGNORE_SERVER_IDS_SYM { $$ = $1 }
| INVOKER_SYM  { $$ = $1 }
| IMPORT { $$ = $1 }
| INDEXES { $$ = $1 }
| INITIAL_SIZE_SYM { $$ = $1 }
| IO_SYM { $$ = $1 }
| IPC_SYM { $$ = $1 }
| ISOLATION { $$ = $1 }
| ISSUER_SYM { $$ = $1 }
| INSERT_METHOD { $$ = $1 }
| KEY_BLOCK_SIZE { $$ = $1 }
| LAST_SYM { $$ = $1 }
| LEAVES { $$ = $1 }
| LESS_SYM { $$ = $1 }
| LEVEL_SYM { $$ = $1 }
| LINESTRING { $$ = $1 }
| LIST_SYM { $$ = $1 }
| LOCAL_SYM { $$ = $1 }
| LOCKS_SYM { $$ = $1 }
| LOGFILE_SYM { $$ = $1 }
| LOGS_SYM { $$ = $1 }
| MAX_ROWS { $$ = $1 }
| MASTER_SYM { $$ = $1 }
| MASTER_HEARTBEAT_PERIOD_SYM { $$ = $1 }
| MASTER_HOST_SYM { $$ = $1 }
| MASTER_PORT_SYM { $$ = $1 }
| MASTER_LOG_FILE_SYM { $$ = $1 }
| MASTER_LOG_POS_SYM { $$ = $1 }
| MASTER_USER_SYM { $$ = $1 }
| MASTER_PASSWORD_SYM { $$ = $1 }
| MASTER_SERVER_ID_SYM { $$ = $1 }
| MASTER_CONNECT_RETRY_SYM { $$ = $1 }
| MASTER_RETRY_COUNT_SYM { $$ = $1 }
| MASTER_DELAY_SYM { $$ = $1 }
| MASTER_SSL_SYM { $$ = $1 }
| MASTER_SSL_CA_SYM { $$ = $1 }
| MASTER_SSL_CAPATH_SYM { $$ = $1 }
| MASTER_SSL_CERT_SYM { $$ = $1 }
| MASTER_SSL_CIPHER_SYM { $$ = $1 }
| MASTER_SSL_CRL_SYM { $$ = $1 }
| MASTER_SSL_CRLPATH_SYM { $$ = $1 }
| MASTER_SSL_KEY_SYM { $$ = $1 }
| MASTER_AUTO_POSITION_SYM { $$ = $1 }
| MAX_CONNECTIONS_PER_HOUR { $$ = $1 }
| MAX_QUERIES_PER_HOUR { $$ = $1 }
| MAX_SIZE_SYM { $$ = $1 }
| MAX_UPDATES_PER_HOUR { $$ = $1 }
| MAX_USER_CONNECTIONS_SYM { $$ = $1 }
| MEDIUM_SYM { $$ = $1 }
| MEMORY_SYM { $$ = $1 }
| MERGE_SYM { $$ = $1 }
| MESSAGE_TEXT_SYM { $$ = $1 }
| MICROSECOND_SYM { $$ = $1 }
| MIGRATE_SYM { $$ = $1 }
| MINUTE_SYM { $$ = $1 }
| MIN_ROWS { $$ = $1 }
| MODIFY_SYM { $$ = $1 }
| MODE_SYM { $$ = $1 }
| MONTH_SYM { $$ = $1 }
| MULTILINESTRING { $$ = $1 }
| MULTIPOINT { $$ = $1 }
| MULTIPOLYGON { $$ = $1 }
| MUTEX_SYM { $$ = $1 }
| MYSQL_ERRNO_SYM { $$ = $1 }
| NAME_SYM { $$ = $1 }
| NAMES_SYM { $$ = $1 }
| NATIONAL_SYM { $$ = $1 }
| NCHAR_SYM { $$ = $1 }
| NDBCLUSTER_SYM { $$ = $1 }
| NEXT_SYM { $$ = $1 }
| NEW_SYM { $$ = $1 }
| NO_WAIT_SYM { $$ = $1 }
| NODEGROUP_SYM { $$ = $1 }
| NONE_SYM { $$ = $1 }
| NUMBER_SYM { $$ = $1 }
| NVARCHAR_SYM { $$ = $1 }
| OFFSET_SYM { $$ = $1 }
| OLD_PASSWORD { $$ = $1 }
| ONE_SYM { $$ = $1 }
| ONLY_SYM { $$ = $1 }
| PACK_KEYS_SYM { $$ = $1 }
| PAGE_SYM { $$ = $1 }
| PARTIAL { $$ = $1 }
| PARTITIONING_SYM { $$ = $1 }
| PARTITIONS_SYM { $$ = $1 }
| PASSWORD { $$ = $1 }
| PHASE_SYM { $$ = $1 }
| PLUGIN_DIR_SYM { $$ = $1 }
| PLUGIN_SYM { $$ = $1 }
| PLUGINS_SYM { $$ = $1 }
| POINT_SYM { $$ = $1 }
| POLYGON { $$ = $1 }
| PRESERVE_SYM { $$ = $1 }
| PREV_SYM { $$ = $1 }
| PRIVILEGES { $$ = $1 }
| PROCESS { $$ = $1 }
| PROCESSLIST_SYM { $$ = $1 }
| PROFILE_SYM { $$ = $1 }
| PROFILES_SYM { $$ = $1 }
| PROXY_SYM { $$ = $1 }
| QUARTER_SYM { $$ = $1 }
| QUERY_SYM { $$ = $1 }
| QUICK { $$ = $1 }
| READ_ONLY_SYM { $$ = $1 }
| REBUILD_SYM { $$ = $1 }
| RECOVER_SYM { $$ = $1 }
| REDO_BUFFER_SIZE_SYM { $$ = $1 }
| REDOFILE_SYM { $$ = $1 }
| REDUNDANT_SYM { $$ = $1 }
| RELAY { $$ = $1 }
| RELAYLOG_SYM { $$ = $1 }
| RELAY_LOG_FILE_SYM { $$ = $1 }
| RELAY_LOG_POS_SYM { $$ = $1 }
| RELAY_THREAD { $$ = $1 }
| RELOAD { $$ = $1 }
| REORGANIZE_SYM { $$ = $1 }
| REPEATABLE_SYM { $$ = $1 }
| REPLICATION { $$ = $1 }
| RESOURCES { $$ = $1 }
| RESUME_SYM { $$ = $1 }
| RETURNED_SQLSTATE_SYM { $$ = $1 }
| RETURNS_SYM { $$ = $1 }
| REVERSE_SYM { $$ = $1 }
| ROLLUP_SYM { $$ = $1 }
| ROUTINE_SYM { $$ = $1 }
| ROWS_SYM { $$ = $1 }
| ROW_COUNT_SYM { $$ = $1 }
| ROW_FORMAT_SYM { $$ = $1 }
| ROW_SYM { $$ = $1 }
| RTREE_SYM { $$ = $1 }
| SCHEDULE_SYM { $$ = $1 }
| SCHEMA_NAME_SYM { $$ = $1 }
| SECOND_SYM { $$ = $1 }
| SERIAL_SYM { $$ = $1 }
| SERIALIZABLE_SYM { $$ = $1 }
| SESSION_SYM { $$ = $1 }
| SIMPLE_SYM { $$ = $1 }
| SHARE_SYM { $$ = $1 }
| SHUTDOWN { $$ = $1 }
| SLOW { $$ = $1 }
| SNAPSHOT_SYM { $$ = $1 }
| SOUNDS_SYM { $$ = $1 }
| SOURCE_SYM { $$ = $1 }
| SQL_AFTER_GTIDS { $$ = $1 }
| SQL_AFTER_MTS_GAPS { $$ = $1 }
| SQL_BEFORE_GTIDS { $$ = $1 }
| SQL_CACHE_SYM { $$ = $1 }
| SQL_BUFFER_RESULT { $$ = $1 }
| SQL_NO_CACHE_SYM { $$ = $1 }
| SQL_THREAD { $$ = $1 }
| STARTS_SYM { $$ = $1 }
| STATS_AUTO_RECALC_SYM { $$ = $1 }
| STATS_PERSISTENT_SYM { $$ = $1 }
| STATS_SAMPLE_PAGES_SYM { $$ = $1 }
| STATUS_SYM { $$ = $1 }
| STORAGE_SYM { $$ = $1 }
| STRING_SYM { $$ = $1 }
| SUBCLASS_ORIGIN_SYM { $$ = $1 }
| SUBDATE_SYM { $$ = $1 }
| SUBJECT_SYM { $$ = $1 }
| SUBPARTITION_SYM { $$ = $1 }
| SUBPARTITIONS_SYM { $$ = $1 }
| SUPER_SYM { $$ = $1 }
| SUSPEND_SYM { $$ = $1 }
| SWAPS_SYM { $$ = $1 }
| SWITCHES_SYM { $$ = $1 }
| TABLE_NAME_SYM { $$ = $1 }
| TABLES { $$ = $1 }
| TABLE_CHECKSUM_SYM { $$ = $1 }
| TABLESPACE { $$ = $1 }
| TEMPORARY { $$ = $1 }
| TEMPTABLE_SYM { $$ = $1 }
| TEXT_SYM { $$ = $1 }
| THAN_SYM { $$ = $1 }
| TRANSACTION_SYM { $$ = $1 }
| TRIGGERS_SYM { $$ = $1 }
| TIMESTAMP { $$ = $1 }
| TIMESTAMP_ADD { $$ = $1 }
| TIMESTAMP_DIFF { $$ = $1 }
| TIME_SYM { $$ = $1 }
| TYPES_SYM { $$ = $1 }
| TYPE_SYM { $$ = $1 }
| UDF_RETURNS_SYM { $$ = $1 }
| FUNCTION_SYM { $$ = $1 }
| UNCOMMITTED_SYM { $$ = $1 }
| UNDEFINED_SYM { $$ = $1 }
| UNDO_BUFFER_SIZE_SYM { $$ = $1 }
| UNDOFILE_SYM { $$ = $1 }
| UNKNOWN_SYM { $$ = $1 }
| UNTIL_SYM { $$ = $1 }
| USER { $$ = $1 }
| USE_FRM { $$ = $1 }
| VARIABLES { $$ = $1 }
| VIEW_SYM { $$ = $1 } 
| VALUE_SYM { $$ = $1 }
| WARNINGS { $$ = $1 } 
| WAIT_SYM { $$ = $1 }
| WEEK_SYM { $$ = $1 }
| WORK_SYM { $$ = $1 }
| WEIGHT_STRING_SYM { $$ = $1 }
| X509_SYM { $$ = $1 }
| XML_SYM { $$ = $1 }
| YEAR_SYM { $$ = $1 }
;

set:
  SET start_option_value_list { $$ = &Set{} };

start_option_value_list:
  option_value_no_option_type option_value_list_continued
| TRANSACTION_SYM transaction_characteristics
| option_type start_option_value_list_following_option_type;

start_option_value_list_following_option_type:
  option_value_following_option_type option_value_list_continued
| TRANSACTION_SYM transaction_characteristics;

option_value_list_continued:
 
| ',' option_value_list;

option_value_list:
  option_value
| option_value_list ',' option_value;

option_value:
  option_type option_value_following_option_type
| option_value_no_option_type;

option_type:
  GLOBAL_SYM
| LOCAL_SYM
| SESSION_SYM;

opt_var_type:
 
| GLOBAL_SYM
| LOCAL_SYM
| SESSION_SYM;

opt_var_ident_type:
 
| GLOBAL_SYM '.'
| LOCAL_SYM '.'
| SESSION_SYM '.';

option_value_following_option_type:
  internal_variable_name equal set_expr_or_default;

option_value_no_option_type:
  internal_variable_name equal set_expr_or_default
| '@' ident_or_text equal expr
| '@' '@' opt_var_ident_type internal_variable_name equal set_expr_or_default
| charset old_or_new_charset_name_or_default
| NAMES_SYM equal expr
| NAMES_SYM charset_name_or_default opt_collate
| PASSWORD equal text_or_password
| PASSWORD FOR_SYM user equal text_or_password;

internal_variable_name:
  ident
| ident '.' ident
| DEFAULT '.' ident;

transaction_characteristics:
  transaction_access_mode
| isolation_level
| transaction_access_mode ',' isolation_level
| isolation_level ',' transaction_access_mode;

transaction_access_mode:
  transaction_access_mode_types;

isolation_level:
  ISOLATION LEVEL_SYM isolation_types;

transaction_access_mode_types:
  READ_SYM ONLY_SYM
| READ_SYM WRITE_SYM;

isolation_types:
  READ_SYM UNCOMMITTED_SYM
| READ_SYM COMMITTED_SYM
| REPEATABLE_SYM READ_SYM
| SERIALIZABLE_SYM;

text_or_password:
  TEXT_STRING
| PASSWORD '(' TEXT_STRING ')'
| OLD_PASSWORD '(' TEXT_STRING ')';

set_expr_or_default:
  expr
| DEFAULT
| ON
| ALL
| BINARY;

lock:
  LOCK_SYM table_or_tables table_lock_list { $$ = &Lock{} };

table_or_tables:
  TABLE_SYM
| TABLES;

table_lock_list:
  table_lock
| table_lock_list ',' table_lock;

table_lock:
  table_ident opt_table_alias lock_option;

lock_option:
  READ_SYM
| WRITE_SYM
| LOW_PRIORITY WRITE_SYM
| READ_SYM LOCAL_SYM;

unlock:
  UNLOCK_SYM table_or_tables { $$ = &Unlock{} };

handler:
  HANDLER_SYM table_ident OPEN_SYM opt_table_alias { $$ = &Handler{} }
| HANDLER_SYM table_ident_nodb CLOSE_SYM { $$ = &Handler{} }
| HANDLER_SYM table_ident_nodb READ_SYM handler_read_or_scan where_clause opt_limit_clause { $$ = &Handler{} }

;

handler_read_or_scan:
  handler_scan_function
| ident handler_rkey_function;

handler_scan_function:
  FIRST_SYM
| NEXT_SYM;

handler_rkey_function:
  FIRST_SYM
| NEXT_SYM
| PREV_SYM
| LAST_SYM
| handler_rkey_mode '(' values ')';

handler_rkey_mode:
  EQ
| GE
| LE
| GT_SYM
| LT;

revoke:
  REVOKE clear_privileges revoke_command { $$ = &Revoke{} };

revoke_command:
  grant_privileges ON opt_table grant_ident FROM grant_list
| grant_privileges ON FUNCTION_SYM grant_ident FROM grant_list
| grant_privileges ON PROCEDURE_SYM grant_ident FROM grant_list
| ALL opt_privileges ',' GRANT OPTION FROM grant_list
| PROXY_SYM ON user FROM grant_list;

grant:
  GRANT clear_privileges grant_command { $$ = &Grant{} };

grant_command:
  grant_privileges ON opt_table grant_ident TO_SYM grant_list require_clause grant_options
| grant_privileges ON FUNCTION_SYM grant_ident TO_SYM grant_list require_clause grant_options
| grant_privileges ON PROCEDURE_SYM grant_ident TO_SYM grant_list require_clause grant_options
| PROXY_SYM ON user TO_SYM grant_list opt_grant_option;

opt_table:
 
| TABLE_SYM;

grant_privileges:
  object_privilege_list
| ALL opt_privileges;

opt_privileges:
 
| PRIVILEGES;

object_privilege_list:
  object_privilege
| object_privilege_list ',' object_privilege;

object_privilege:
  SELECT_SYM opt_column_list
| INSERT opt_column_list
| UPDATE_SYM opt_column_list
| REFERENCES opt_column_list
| DELETE_SYM
| USAGE
| INDEX_SYM
| ALTER
| CREATE
| DROP
| EXECUTE_SYM
| RELOAD
| SHUTDOWN
| PROCESS
| FILE_SYM
| GRANT OPTION
| SHOW DATABASES
| SUPER_SYM
| CREATE TEMPORARY TABLES
| LOCK_SYM TABLES
| REPLICATION SLAVE
| REPLICATION CLIENT_SYM
| CREATE VIEW_SYM
| SHOW VIEW_SYM
| CREATE ROUTINE_SYM
| ALTER ROUTINE_SYM
| CREATE USER
| EVENT_SYM
| TRIGGER_SYM
| CREATE TABLESPACE;

opt_and:
 
| AND_SYM;

require_list:
  require_list_element opt_and require_list
| require_list_element;

require_list_element:
  SUBJECT_SYM TEXT_STRING
| ISSUER_SYM TEXT_STRING
| CIPHER_SYM TEXT_STRING;

grant_ident:
  '*'
| ident '.' '*'
| '*' '.' '*'
| table_ident;

user_list:
  user
| user_list ',' user;

grant_list:
  grant_user
| grant_list ',' grant_user;

grant_user:
  user IDENTIFIED_SYM BY TEXT_STRING
| user IDENTIFIED_SYM BY PASSWORD TEXT_STRING
| user IDENTIFIED_SYM WITH ident_or_text
| user IDENTIFIED_SYM WITH ident_or_text AS TEXT_STRING_sys
| user;

opt_column_list:
 
| '(' column_list ')';

column_list:
  column_list ',' column_list_id
| column_list_id;

column_list_id:
  ident;

require_clause:
 
| REQUIRE_SYM require_list
| REQUIRE_SYM SSL_SYM
| REQUIRE_SYM X509_SYM
| REQUIRE_SYM NONE_SYM;

grant_options:
 
| WITH grant_option_list;

opt_grant_option:
 
| WITH GRANT OPTION;

grant_option_list:
  grant_option_list grant_option
| grant_option;

grant_option:
  GRANT OPTION
| MAX_QUERIES_PER_HOUR ulong_num
| MAX_UPDATES_PER_HOUR ulong_num
| MAX_CONNECTIONS_PER_HOUR ulong_num
| MAX_USER_CONNECTIONS_SYM ulong_num;

begin:
  BEGIN_SYM opt_work;

opt_work:
 
| WORK_SYM;

opt_chain:
 
| AND_SYM NO_SYM CHAIN_SYM
| AND_SYM CHAIN_SYM;

opt_release:
 
| RELEASE_SYM
| NO_SYM RELEASE_SYM;

opt_savepoint:
 
| SAVEPOINT_SYM;

commit:
  COMMIT_SYM opt_work opt_chain opt_release { $$ = &Commit{} };

rollback:
  ROLLBACK_SYM opt_work opt_chain opt_release { $$ = &Rollback{} }
| ROLLBACK_SYM opt_work TO_SYM opt_savepoint ident { $$ = &Rollback{} }
;

savepoint:
  SAVEPOINT_SYM ident { $$ = &SavePoint{} }
;

release:
  RELEASE_SYM SAVEPOINT_SYM ident { $$ = &Release{} };

union_clause:
 
| union_list;

union_list:
  UNION_SYM union_option select_init;

union_opt:
 
| union_list
| union_order_or_limit;

opt_union_order_or_limit:
 
| union_order_or_limit;

union_order_or_limit:
  order_or_limit;

order_or_limit:
  order_clause opt_limit_clause_init
| limit_clause;

union_option:
 
| DISTINCT
| ALL;

query_specification:
  SELECT_SYM select_init2_derived
| '(' select_paren_derived ')';

query_expression_body:
  query_specification opt_union_order_or_limit
| query_expression_body UNION_SYM union_option query_specification opt_union_order_or_limit;

subselect:
  subselect_start query_expression_body subselect_end;

subselect_start:
 ;

subselect_end:
 ;

opt_query_expression_options:
 
| query_expression_option_list;

query_expression_option_list:
  query_expression_option_list query_expression_option
| query_expression_option;

query_expression_option:
  STRAIGHT_JOIN
| HIGH_PRIORITY
| DISTINCT
| SQL_SMALL_RESULT
| SQL_BIG_RESULT
| SQL_BUFFER_RESULT
| SQL_CALC_FOUND_ROWS
| ALL;

view_or_trigger_or_sp_or_event:
  definer definer_tail
| no_definer no_definer_tail
| view_replace_or_algorithm definer_opt view_tail;

definer_tail:
  view_tail
| trigger_tail
| sp_tail
| sf_tail
| event_tail;

no_definer_tail:
  view_tail
| trigger_tail
| sp_tail
| sf_tail
| udf_tail
| event_tail;

definer_opt:
  no_definer
| definer;

no_definer:
 ;

definer:
  DEFINER_SYM EQ user;

view_replace_or_algorithm:
  view_replace
| view_replace view_algorithm
| view_algorithm;

view_replace:
  OR_SYM REPLACE;

view_algorithm:
  ALGORITHM_SYM EQ UNDEFINED_SYM
| ALGORITHM_SYM EQ MERGE_SYM
| ALGORITHM_SYM EQ TEMPTABLE_SYM;

view_suid:
 
| SQL_SYM SECURITY_SYM DEFINER_SYM
| SQL_SYM SECURITY_SYM INVOKER_SYM;

view_tail:
  view_suid VIEW_SYM table_ident view_list_opt AS view_select;

view_list_opt:
 
| '(' view_list ')';

view_list:
  ident
| view_list ',' ident;

view_select:
  view_select_aux view_check_option;

view_select_aux:
  create_view_select union_clause
| '(' create_view_select_paren ')' union_opt;

create_view_select_paren:
  create_view_select
| '(' create_view_select_paren ')';

create_view_select:
  SELECT_SYM select_part2;

view_check_option:
 
| WITH CHECK_SYM OPTION
| WITH CASCADED CHECK_SYM OPTION
| WITH LOCAL_SYM CHECK_SYM OPTION;

trigger_tail:
  TRIGGER_SYM remember_name sp_name trg_action_time trg_event ON remember_name table_ident FOR_SYM remember_name EACH_SYM ROW_SYM sp_proc_stmt;

udf_tail:
  AGGREGATE_SYM remember_name FUNCTION_SYM ident RETURNS_SYM udf_type SONAME_SYM TEXT_STRING_sys
| remember_name FUNCTION_SYM ident RETURNS_SYM udf_type SONAME_SYM TEXT_STRING_sys;

sf_tail:
  remember_name FUNCTION_SYM sp_name '(' sp_fdparam_list ')' RETURNS_SYM type_with_opt_collate sp_c_chistics sp_proc_stmt;

sp_tail:
  PROCEDURE_SYM remember_name sp_name '(' sp_pdparam_list ')' sp_c_chistics sp_proc_stmt;

xa:
  XA_SYM begin_or_start xid opt_join_or_resume { $$ = &XA{} }
| XA_SYM END xid opt_suspend { $$ = &XA{} }
| XA_SYM PREPARE_SYM xid { $$ = &XA{} }
| XA_SYM COMMIT_SYM xid opt_one_phase { $$ = &XA{} }
| XA_SYM ROLLBACK_SYM xid { $$ = &XA{} }
| XA_SYM RECOVER_SYM { $$ = &XA{} }
;

xid:
  text_string
| text_string ',' text_string
| text_string ',' text_string ',' ulong_num;

begin_or_start:
  BEGIN_SYM
| START_SYM;

opt_join_or_resume:
 
| JOIN_SYM
| RESUME_SYM;

opt_one_phase:
 
| ONE_SYM PHASE_SYM;

opt_suspend:
 
| SUSPEND_SYM opt_migrate;

opt_migrate:
 
| FOR_SYM MIGRATE_SYM;

install:
  INSTALL_SYM PLUGIN_SYM ident SONAME_SYM TEXT_STRING_sys { $$ = &Install{} }
;

uninstall:
  UNINSTALL_SYM PLUGIN_SYM ident { $$ = &Uninstall{} }
;
%%
