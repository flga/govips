package vips

/*
#cgo pkg-config: vips
#include "bridge.h"

extern long govips_obs_write(VipsTarget*, const void*, long, void*);
extern void govips_obs_finish(VipsTarget*, void*);

static inline long cgovips_obs_write(VipsTarget *target, const void *buffer, long length, void *user) {
	return govips_obs_write(target, buffer, length, user);
}

static inline void cgovips_obs_finish(VipsTarget *target, void *user) {
	govips_obs_finish(target, user);
}

VipsTarget* govips_new_writer_target(void* user) {
	VipsTargetCustom *target = vips_target_custom_new();
	g_signal_connect(target, "write", G_CALLBACK(cgovips_obs_write), user);
	g_signal_connect(target, "finish", G_CALLBACK(cgovips_obs_finish), user);
	return VIPS_TARGET(target);
}

*/
import "C"
