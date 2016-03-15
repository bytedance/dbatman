// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

const (
	MinProtocolVersion byte = minProtocolVersion
	MaxPacketSize           = maxPacketSize
	TimeFormat              = timeFormat
)

// MySQL constants documentation:
// http://dev.mysql.com/doc/internals/en/client-server-protocol.html

const (
	OK          byte = iOK
	LocalInFile byte = iLocalInFile
	EOF         byte = iEOF
	ERR         byte = iERR
)

const (
	ClientLongPassword               clientFlag = clientLongPassword
	ClientFoundRows                  clientFlag = clientFoundRows
	ClientLongFlag                   clientFlag = clientLongFlag
	ClientConnectWithDB              clientFlag = clientConnectWithDB
	ClientNoSchema                   clientFlag = clientNoSchema
	ClientCompress                   clientFlag = clientCompress
	ClientODBC                       clientFlag = clientODBC
	ClientLocalFiles                 clientFlag = clientLocalFiles
	ClientIgnoreSpace                clientFlag = clientIgnoreSpace
	ClientProtocol                   clientFlag = clientProtocol41
	ClientInteractive                clientFlag = clientInteractive
	ClientSSL                        clientFlag = clientSSL
	ClientIgnoreSIGPIPE              clientFlag = clientIgnoreSIGPIPE
	ClientTransactions               clientFlag = clientTransactions
	ClientReserved                   clientFlag = clientReserved
	ClientSecureConn                 clientFlag = clientSecureConn
	ClientMultiStatements            clientFlag = clientMultiStatements
	ClientMultiResults               clientFlag = clientMultiResults
	ClientPSMultiResults             clientFlag = clientPSMultiResults
	ClientPluginAuth                 clientFlag = clientPluginAuth
	ClientConnectAttrs               clientFlag = clientConnectAttrs
	ClientPluginAuthLenEncClientData clientFlag = clientPluginAuthLenEncClientData
	ClientCanHandleExpiredPasswords  clientFlag = clientCanHandleExpiredPasswords
	ClientSessionTrack               clientFlag = clientSessionTrack
	ClientDeprecateEOF               clientFlag = clientDeprecateEOF
)

const (
	ComQuit             byte = comQuit
	ComInitDB           byte = comInitDB
	ComQuery            byte = comQuery
	ComFieldList        byte = comFieldList
	ComCreateDB         byte = comCreateDB
	ComDropDB           byte = comDropDB
	ComRefresh          byte = comRefresh
	ComShutdown         byte = comShutdown
	ComStatistics       byte = comStatistics
	ComProcessInfo      byte = comProcessInfo
	ComConnect          byte = comConnect
	ComProcessKill      byte = comProcessKill
	ComDebug            byte = comDebug
	ComPing             byte = comPing
	ComTime             byte = comTime
	ComDelayedInsert    byte = comDelayedInsert
	ComChangeUser       byte = comChangeUser
	ComBinlogDump       byte = comBinlogDump
	ComTableDump        byte = comTableDump
	ComConnectOut       byte = comConnectOut
	ComRegisterSlave    byte = comRegisterSlave
	ComStmtPrepare      byte = comStmtPrepare
	ComStmtExecute      byte = comStmtExecute
	ComStmtSendLongData byte = comStmtSendLongData
	ComStmtClose        byte = comStmtClose
	ComStmtReset        byte = comStmtReset
	ComSetOption        byte = comSetOption
	ComStmtFetch        byte = comStmtFetch
)

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType
const (
	FieldTypeDecimal   byte = fieldTypeDecimal
	FieldTypeTiny      byte = fieldTypeTiny
	FieldTypeShort     byte = fieldTypeShort
	FieldTypeLong      byte = fieldTypeLong
	FieldTypeFloat     byte = fieldTypeFloat
	FieldTypeDouble    byte = fieldTypeDouble
	FieldTypeNULL      byte = fieldTypeNULL
	FieldTypeTimestamp byte = fieldTypeTimestamp
	FieldTypeLongLong  byte = fieldTypeLongLong
	FieldTypeInt       byte = fieldTypeInt24
	FieldTypeDate      byte = fieldTypeDate
	FieldTypeTime      byte = fieldTypeTime
	FieldTypeDateTime  byte = fieldTypeDateTime
	FieldTypeYear      byte = fieldTypeYear
	FieldTypeNewDate   byte = fieldTypeNewDate
	FieldTypeVarChar   byte = fieldTypeVarChar
	FieldTypeBit       byte = fieldTypeBit
)
const (
	FieldTypeJSON       byte = fieldTypeJSON
	FieldTypeNewDecimal byte = fieldTypeNewDecimal
	FieldTypeEnum       byte = fieldTypeEnum
	FieldTypeSet        byte = fieldTypeSet
	FieldTypeTinyBLOB   byte = fieldTypeTinyBLOB
	FieldTypeMediumBLOB byte = fieldTypeMediumBLOB
	FieldTypeLongBLOB   byte = fieldTypeLongBLOB
	FieldTypeBLOB       byte = fieldTypeBLOB
	FieldTypeVarString  byte = fieldTypeVarString
	FieldTypeString     byte = fieldTypeString
	FieldTypeGeometry   byte = fieldTypeGeometry
)

const (
	FlagNotNULL       fieldFlag = flagNotNULL
	FlagPriKey        fieldFlag = flagPriKey
	FlagUniqueKey     fieldFlag = flagUniqueKey
	FlagMultipleKey   fieldFlag = flagMultipleKey
	FlagBLOB          fieldFlag = flagBLOB
	FlagUnsigned      fieldFlag = flagUnsigned
	FlagZeroFill      fieldFlag = flagZeroFill
	FlagBinary        fieldFlag = flagBinary
	FlagEnum          fieldFlag = flagEnum
	FlagAutoIncrement fieldFlag = flagAutoIncrement
	FlagTimestamp     fieldFlag = flagTimestamp
	FlagSet           fieldFlag = flagSet
	FlagUnknown1      fieldFlag = flagUnknown1
	FlagUnknown2      fieldFlag = flagUnknown2
	FlagUnknown3      fieldFlag = flagUnknown3
	FlagUnknown4      fieldFlag = flagUnknown4
)

// http://dev.mysql.com/doc/internals/en/status-flags.html

const (
	StatusInTrans             statusFlag = statusInTrans
	StatusInAutocommit        statusFlag = statusInAutocommit
	StatusReserved            statusFlag = statusReserved // Not in documentation
	StatusMoreResultsExists   statusFlag = statusMoreResultsExists
	StatusNoGoodIndexUsed     statusFlag = statusNoGoodIndexUsed
	StatusNoIndexUsed         statusFlag = statusNoIndexUsed
	StatusCursorExists        statusFlag = statusCursorExists
	StatusLastRowSent         statusFlag = statusLastRowSent
	StatusDbDropped           statusFlag = statusDbDropped
	StatusNoBackslashEscapes  statusFlag = statusNoBackslashEscapes
	StatusMetadataChanged     statusFlag = statusMetadataChanged
	StatusQueryWasSlow        statusFlag = statusQueryWasSlow
	StatusPsOutParams         statusFlag = statusPsOutParams
	StatusInTransReadonly     statusFlag = statusInTransReadonly
	StatusSessionStateChanged statusFlag = statusSessionStateChanged
)
