# Copyright 2012, Google Inc. All rights reserved.
# Use of this source code is governed by a BSD-style license that can
# be found in the LICENSE file.

# MAKEFLAGS = -s

sql.go: sql_yacc.yy
	./bin/yacc -o sql_yacc.go -p MySQL sql_yacc.yy
	gofmt -w sql_yacc.go

clean:
	rm -f y.output sql_yacc.go
