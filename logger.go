/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * Zymatik Genobase - A Human Genomics reference database.
 * Copyright (C) 2024 Damian Peckett <damian@pecke.tt>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
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
