/* SPDX-License-Identifier: MPL-2.0
 *
 * Zymatik Genobase - A human genomics reference DB.
 * Copyright (C) 2024 Damian Peckett <damian@pecke.tt>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the Mozilla Public License v2.0.
 *
 * You should have received a copy of the Mozilla Public License v2.0
 * along with this program. If not, see <https://mozilla.org/MPL/2.0/>.
 */

package genobase

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type gooseLogger struct {
	*slog.Logger
}

func (l *gooseLogger) Printf(format string, v ...any) {
	l.Logger.Info(strings.TrimSpace(fmt.Sprintf(format, v...)))

}

func (l *gooseLogger) Fatalf(format string, v ...any) {
	l.Logger.Error(strings.TrimSpace(fmt.Sprintf(format, v...)))
	os.Exit(1)
}
