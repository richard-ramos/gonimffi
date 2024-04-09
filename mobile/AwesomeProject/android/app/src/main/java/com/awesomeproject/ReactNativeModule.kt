package com.awesomeproject
import com.facebook.react.bridge.NativeModule
import com.facebook.react.bridge.ReactApplicationContext
import com.facebook.react.bridge.ReactContext
import com.facebook.react.bridge.ReactContextBaseJavaModule
import com.facebook.react.bridge.ReactMethod
import com.facebook.react.bridge.Promise
import android.util.Log

class HelloModule(reactContext: ReactApplicationContext) : ReactContextBaseJavaModule(reactContext) {
    override fun getName() = "HelloModule"

    external fun helloWorldJNI(): String?

    @ReactMethod
    fun start(promise: Promise) {
        var a  = helloWorldJNI()
        Log.d("HelloModule", "$a")
        promise.resolve(a);
    }
}