package msmnile

import (
    "android/soong/android"
)

func init() {
    android.RegisterModuleType("meizu_msmnile_fod_hal_binary", fodHalBinaryFactory)
    android.RegisterModuleType("meizu_msmnile_light_hal_binary", lightHalBinaryFactory)
}
