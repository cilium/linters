// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package xexp

import (
	"context"
	"io"

	"golang.org/x/exp/slog"
)

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	// calling any of the logging function shouldn't be flagged
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.DebugContext(ctx, "debug")
	log.InfoContext(ctx, "info")
	log.WarnContext(ctx, "warn")
	log.ErrorContext(ctx, "error")
	log.DebugCtx(ctx, "debug")
	log.InfoCtx(ctx, "info")
	log.WarnCtx(ctx, "warn")
	log.ErrorCtx(ctx, "error")
	log.Log(ctx, slog.LevelInfo, "log")
	log.LogAttrs(ctx, slog.LevelInfo, "log")

	// creating a new logger with With or WithGroup for use later shouldn't
	// be flagged
	wlog := log.With("key", "va")
	wlog.Info("info")
	glog := log.WithGroup("group")
	glog.Info("info")

	// With
	log.With("key", "val").Debug("debug")                       // want `call to Debug on a newly instantiated Logger`
	log.With("key", "val").Info("info")                         // want `call to Info on a newly instantiated Logger`
	log.With("key", "val").Warn("warn")                         // want `call to Warn on a newly instantiated Logger`
	log.With("key", "val").Error("error")                       // want `call to Error on a newly instantiated Logger`
	log.With("key", "val").DebugContext(ctx, "debug")           // want `call to DebugContext on a newly instantiated Logger`
	log.With("key", "val").InfoContext(ctx, "info")             // want `call to InfoContext on a newly instantiated Logger`
	log.With("key", "val").WarnContext(ctx, "warn")             // want `call to WarnContext on a newly instantiated Logger`
	log.With("key", "val").ErrorContext(ctx, "error")           // want `call to ErrorContext on a newly instantiated Logger`
	log.With("key", "val").DebugCtx(ctx, "debug")               // want `call to DebugCtx on a newly instantiated Logger`
	log.With("key", "val").InfoCtx(ctx, "info")                 // want `call to InfoCtx on a newly instantiated Logger`
	log.With("key", "val").WarnCtx(ctx, "warn")                 // want `call to WarnCtx on a newly instantiated Logger`
	log.With("key", "val").ErrorCtx(ctx, "error")               // want `call to ErrorCtx on a newly instantiated Logger`
	log.With("key", "val").Log(ctx, slog.LevelInfo, "log")      // want `call to Log on a newly instantiated Logger`
	log.With("key", "val").LogAttrs(ctx, slog.LevelInfo, "log") // want `call to LogAttrs on a newly instantiated Logger`

	// WithGroup
	log.WithGroup("group").Debug("debug")                       // want `call to Debug on a newly instantiated Logger`
	log.WithGroup("group").Info("info")                         // want `call to Info on a newly instantiated Logger`
	log.WithGroup("group").Warn("warn")                         // want `call to Warn on a newly instantiated Logger`
	log.WithGroup("group").Error("error")                       // want `call to Error on a newly instantiated Logger`
	log.WithGroup("group").DebugContext(ctx, "debug")           // want `call to DebugContext on a newly instantiated Logger`
	log.WithGroup("group").InfoContext(ctx, "info")             // want `call to InfoContext on a newly instantiated Logger`
	log.WithGroup("group").WarnContext(ctx, "warn")             // want `call to WarnContext on a newly instantiated Logger`
	log.WithGroup("group").ErrorContext(ctx, "error")           // want `call to ErrorContext on a newly instantiated Logger`
	log.WithGroup("group").DebugCtx(ctx, "debug")               // want `call to DebugCtx on a newly instantiated Logger`
	log.WithGroup("group").InfoCtx(ctx, "info")                 // want `call to InfoCtx on a newly instantiated Logger`
	log.WithGroup("group").WarnCtx(ctx, "warn")                 // want `call to WarnCtx on a newly instantiated Logger`
	log.WithGroup("group").ErrorCtx(ctx, "error")               // want `call to ErrorCtx on a newly instantiated Logger`
	log.WithGroup("group").Log(ctx, slog.LevelInfo, "log")      // want `call to Log on a newly instantiated Logger`
	log.WithGroup("group").LogAttrs(ctx, slog.LevelInfo, "log") // want `call to LogAttrs on a newly instantiated Logger`
}
