// Copyright (c) All Rights Reserved
// @file    package.go
// @author  王靖 (wangjild@gmail.com)
// @date    14-11-27 14:40:10
// @version $Revision: 1.0 $
// @brief

package mysql

import ()

// Packets documentation:
// http://dev.mysql.com/doc/internals/en/mysql-packet.html
const MaxPacketSize = 1<<24 - 1
const PacketHeadSize = 4

/* vim: set expandtab ts=4 sw=4 */
