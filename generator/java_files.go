package generator

const JAVA_MAIN = `package com.huawei.hms.cordova.%s;

import android.content.Intent;

import org.apache.cordova.CordovaWebView;
import org.apache.cordova.CallbackContext;
import org.apache.cordova.CordovaPlugin;
import org.apache.cordova.CordovaInterface;

import org.json.JSONArray;

import com.huawei.hms.cordova.%s.basef.CordovaBaseModule;
import com.huawei.hms.cordova.%s.basef.handler.CordovaController;

import java.util.Arrays;

public class HMS%s extends CordovaPlugin {
    
    private static final String SERVICE = "<service-name>";
    private static final String VERSION = "<version>";
    private CordovaController cordovaController;

    @Override
    public void initialize(CordovaInterface cordova, CordovaWebView webView) {
        super.initialize(cordova, webView);
        cordovaController = new CordovaController(this, SERVICE, VERSION,
                Arrays.asList(new CordovaBaseModule[]{
                        new Test(webView.getContext())
                }));
    }

    @Override
    public boolean execute(String action, JSONArray args, final CallbackContext callbackContext) {
        return cordovaController.execute(action, args, callbackContext);
    }

    @Override
    public void onActivityResult(int requestCode, int resultCode, Intent intent) {
        super.onActivityResult(requestCode, resultCode, intent);
        cordovaController.onActivityResult(requestCode, resultCode, intent);
    }

    @Override
    public void onRequestPermissionResult(int requestCode, String[] permissions, int[] grantResults) {
        super.onRequestPermissionResult(requestCode, permissions, grantResults);
        cordovaController.onRequestPermissionResult(requestCode, permissions, grantResults);
    }

    @Override
    public void onPause(boolean multitasking) {
        super.onPause(multitasking);
        cordovaController.onPause(multitasking);
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        cordovaController.onDestroy();
    }

    @Override
    public void onReset() {
        super.onReset();
        cordovaController.onReset();
    }

    @Override
    public void onResume(boolean multitasking) {
        super.onResume(multitasking);
        cordovaController.onResume(multitasking);
    }

    @Override
    public void onStart() {
        super.onStart();
        cordovaController.onStart();
    }

    @Override
    public void onStop() {
        super.onStop();
        cordovaController.onStop();
    }
}
`

const JAVA_EXAMPLE = `package com.huawei.hms.cordova.%s;

import android.content.Context;
import android.widget.Toast;

import com.huawei.hms.cordova.%s.basef.CordovaBaseModule;
import com.huawei.hms.cordova.%s.basef.CordovaMethod;
import com.huawei.hms.cordova.%s.basef.HMSLog;
import com.huawei.hms.cordova.%s.basef.handler.CorPack;
import com.huawei.hms.cordova.%s.basef.handler.Promise;

import org.json.JSONArray;
import org.json.JSONException;

public class Test extends CordovaBaseModule {
    private Context context;
    public Test(Context context) {
        this.context = context;
    }

    @HMSLog
    @CordovaMethod
    public void showToast(final CorPack corPack, JSONArray args, final Promise promise) throws JSONException {
        String message = args.getString(0);
        Toast.makeText(context, message, Toast.LENGTH_SHORT).show();
        promise.success();
    }
}

`

const JAVAC_BASE_ANNOTATION = `package com.huawei.hms.cordova.%s.basef;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface %s {
}
`

const JAVAC_CORBASE_MODULE = `package com.huawei.hms.cordova.%s.basef;

import android.content.Intent;

public abstract class CordovaBaseModule {
    private final String reference;

    public CordovaBaseModule() {
        this.reference = this.getClass().getSimpleName();
    }
    public void onDestroy(){}
    public void onPause(boolean multitasking){}
    public void onResume(boolean multitasking){}
    public void onReset(){}
    public void onStart(){}
    public void onStop(){}
    public void onActivityResult(int requestCode, int resultCode, Intent data){}
    public void onRequestPermissionResult(int requestCode, String[] permissions, int[] grantResults){}


    public String getReference(){
        return reference;
    }
}
`

const JAVAC_NSCM_EXCEPTION = `package com.huawei.hms.cordova.%s.basef.handler;

public class NoSuchCordovaModuleException extends RuntimeException {
}
`
const JAVAC_PROMISE = `package com.huawei.hms.cordova.%s.basef.handler;

import org.apache.cordova.CallbackContext;
import org.apache.cordova.PluginResult;
import org.json.JSONArray;
import org.json.JSONObject;

import static org.apache.cordova.PluginResult.Status.OK;

public class Promise {

    private final CallbackContext callbackContext;
    private final HMSLogger hmsLogger;
    private final String methodName;
    private final boolean isLoggerRunning;

    public Promise(final CallbackContext callbackContext, final HMSLogger logger, String method, boolean isActive) {
        this.callbackContext = callbackContext;
        this.hmsLogger = logger;
        this.methodName = method;
        this.isLoggerRunning = isActive;
    }

    public void success() {
        callbackContext.success();
        sendLogEvent(null);
    }
    public void success(int message) {
        callbackContext.success(message);
        sendLogEvent(null);
    }
    public void success(byte[] message) {
        callbackContext.success(message);
        sendLogEvent(null);
    }
    public void success(String message) {
        callbackContext.success(message);
        sendLogEvent(null);
    }
    public void success(JSONArray message) {
        callbackContext.success(message);
        sendLogEvent(null);
    }
    public void success(JSONObject message) {
        callbackContext.success(message);
        sendLogEvent(null);
    }
    public void success(boolean message) {
        callbackContext.sendPluginResult(new PluginResult(OK, message));
        sendLogEvent(null);
    }
    public void success(float message) {
        callbackContext.sendPluginResult(new PluginResult(OK, message));
        sendLogEvent(null);
    }
    public void error(int message) {
        callbackContext.error(message);
        sendLogEvent("" + message);
    }
    public void error(String message) {
        callbackContext.error(message);
        sendLogEvent(message);
    }
    public void error(JSONObject message) {
        callbackContext.error(message);
        sendLogEvent(message.toString());
    }
    public void sendPluginResult(PluginResult pluginResult) {
        callbackContext.sendPluginResult(pluginResult);
        sendLogEvent(null);
    }

    private void sendLogEvent(String nullable) {
        if (!isLoggerRunning) return;
        if (nullable == null) hmsLogger.sendSingleEvent(methodName);
        else hmsLogger.sendSingleEvent(methodName, nullable);
    }

}
`
const JAVAC_HMS_LOGGER = `package com.huawei.hms.cordova.%s.basef.handler;

import static android.os.Build.DEVICE;

import android.content.Context;
import android.content.pm.PackageManager;
import android.net.ConnectivityManager;
import android.net.NetworkInfo;
import android.util.Log;

import com.huawei.agconnect.config.AGConnectServicesConfig;
import com.huawei.hms.support.hianalytics.HiAnalyticsUtils;
import com.huawei.hms.utils.HMSBIInitializer;

import java.lang.ref.WeakReference;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

final class HMSLogger {
    private static final String TAG = "HMSLogger";

    private static final String PLATFORM = "Cordova";
    private static final String SERVICE = "Cross-Platform";

    private static final String SUCCESS = "0";
    private static final String UNKNOWN = "UNKNOWN";
    private static final String NOT_AVAILABLE = "NOT_AVAILABLE";

    private static final String SINGLE_EVENT_ID = "60000";
    private static final String PERIODIC_EVENT_ID = "60001";

    private static final String NETWORK_TYPE_WIFI = "WIFI";

    private static volatile HMSLogger instance;

    private final String kit;
    private final String version;

    private final WeakReference<Context> weakContext;
    private final HiAnalyticsUtils hiAnalyticsUtils;
    private final ConnectivityManager connectivityManager;

    private final Map<String, Object> singleEventMap = new HashMap<>();
    private final Map<String, Object> periodicEventMap = new HashMap<>();
    private final Map<String, Long> allCountMap = new HashMap<>();
    private final Map<String, Long> failCountMap = new HashMap<>();
    private final Map<String, Long> startTimeMap = new HashMap<>();
    private final Map<String, Long> firstReceiveTimeMap = new HashMap<>();
    private final Map<String, Long> lastReceiveTimeMap = new HashMap<>();
    private final Map<String, Map<String, Long>> resultCodeCountMap = new HashMap<>();
    private final Map<Integer, String> networkTypeMap = createNetworkTypeMap();

    private boolean isEnabled = false;

    /**
     * Private constructor of this class.
     *
     * @param context Application's context
     */
    private HMSLogger(final Context context, final String kit, final String version) {
        this.kit = kit;
        this.version = version;
        weakContext = new WeakReference<>(context);
        hiAnalyticsUtils = HiAnalyticsUtils.getInstance();
        connectivityManager = objectCast(context.getSystemService(Context.CONNECTIVITY_SERVICE),
            ConnectivityManager.class);

        hiAnalyticsUtils.enableLog();
        HMSBIInitializer.getInstance(context).initBI();
        setupEventMap(singleEventMap);
        setupEventMap(periodicEventMap);
        enableLogger();
    }

    /**
     * Returns the instance of this class.
     *
     * @param context Context object
     * @return HMSLogger instance
     */
    static synchronized HMSLogger getInstance(final Context context, final String kit, final String version) {
        if (instance == null) {
            synchronized (HMSLogger.class) {
                if (instance == null) {
                    instance = new HMSLogger(context.getApplicationContext(), kit, version);
                }
            }
        }
        return instance;
    }

    /**
     * Returns actual context reference.
     *
     * @return Actual context reference
     */
    private synchronized Context getContext() {
        return weakContext.get();
    }

    /**
     * Enables HMSLogger.
     */
    synchronized void enableLogger() {
        isEnabled = true;
        Log.d(TAG, "HMS Plugin Dotting is Enabled!");
    }

    /**
     * Disables HMSLogger.
     */
    synchronized void disableLogger() {
        isEnabled = false;
        Log.d(TAG, "HMS Plugin Dotting is Disabled!");
    }

    /**
     * Sets method start time for given method name.
     *
     * @param methodName Name of the method that will be logged
     */
    synchronized void startMethodExecutionTimer(final String methodName) {
        startTimeMap.put(methodName, System.currentTimeMillis());
    }

    /**
     * Sends successful single event.
     *
     * @param methodName The name of the method called
     */
    synchronized void sendSingleEvent(final String methodName) {
        sendEvent(SINGLE_EVENT_ID, methodName, SUCCESS);
    }

    /**
     * Sends unsuccessful single event.
     *
     * @param methodName The name of the method called.
     * @param errorCode  API error code
     */
    synchronized void sendSingleEvent(final String methodName, final String errorCode) {
        sendEvent(SINGLE_EVENT_ID, methodName, errorCode);
    }

    /**
     * Sends successful periodic event.
     *
     * @param methodName The name of the method called
     */
    synchronized void sendPeriodicEvent(final String methodName) {
        sendEvent(PERIODIC_EVENT_ID, methodName, SUCCESS);
    }

    /**
     * Sends unsuccessful periodic event.
     *
     * @param methodName The name of the method called
     * @param errorCode  API error code
     */
    synchronized void sendPeriodicEvent(final String methodName, final String errorCode) {
        sendEvent(PERIODIC_EVENT_ID, methodName, errorCode);
    }

    /**
     * Sends the event based on eventId, methodName, and resultCode.
     *
     * @param eventId    Constant id of the event
     * @param methodName The name of the method called
     * @param resultCode Code of the method's result. "0" for success, others for error
     */
    private synchronized void sendEvent(final String eventId, final String methodName, final String resultCode) {
        if (isEnabled) {
            final long currentTime = System.currentTimeMillis();

            if (eventId.equals(SINGLE_EVENT_ID)) {
                putToSingleEventMap(methodName, resultCode, currentTime);
                hiAnalyticsUtils.onNewEvent(getContext(), SINGLE_EVENT_ID, singleEventMap);

                Log.d(TAG, "singleEventMap -> " + singleEventMap);
            } else {
                putToPeriodicEventMap(methodName, resultCode, currentTime);
                hiAnalyticsUtils.onNewEvent(getContext(), PERIODIC_EVENT_ID, periodicEventMap);

                Log.d(TAG, "periodicEventMap -> " + periodicEventMap);
            }
        }
    }

    /**
     * Gets "client/app_id" value from agconnect-services.json file.
     *
     * @return app_id value or NOT_AVAILABLE if not found
     */
    private synchronized String getAppId() {
        try {
            return AGConnectServicesConfig.fromContext(getContext()).getString("client/app_id");
        } catch (final NullPointerException e) {
            Log.d(TAG, "AgConnect is not found. Setting appId value to " + NOT_AVAILABLE);
        }
        return NOT_AVAILABLE;
    }

    /**
     * Gets app version name.
     *
     * @param packageName Package name of the app
     * @return App version name in String type
     */
    private synchronized String getAppVersionName(final String packageName) {
        try {
            return getContext().getPackageManager().getPackageInfo(packageName, 0).versionName;
        } catch (final PackageManager.NameNotFoundException e) {
            Log.e(TAG, "getAppVersionName ->  Could not get appVersionName!");
            return NOT_AVAILABLE;
        }
    }

    /**
     * Detects current network type.
     *
     * @return Human readable network type; such as WIFI, 4G
     */
    private synchronized String getNetworkType() {
        if (connectivityManager != null) {
            final NetworkInfo networkInfo = connectivityManager.getActiveNetworkInfo();
            if (networkInfo != null && networkInfo.isConnected()) {
                final int networkType = networkInfo.getType();
                if (ConnectivityManager.TYPE_WIFI == networkType) {
                    return NETWORK_TYPE_WIFI;
                } else if (ConnectivityManager.TYPE_MOBILE == networkType) {
                    final int networkSubType = networkInfo.getSubtype();
                    return getOrDefault(networkTypeMap, networkSubType, UNKNOWN);
                } else {
                    return UNKNOWN;
                }
            } else {
                return NOT_AVAILABLE;
            }
        } else {
            return NOT_AVAILABLE;
        }
    }

    /**
     * Sets default values to given map.
     *
     * @param map HashMap to put default values
     */
    private synchronized void setupEventMap(final Map<String, Object> map) {
        map.put("kit", kit);
        map.put("platform", PLATFORM);
        map.put("version", version);
        map.put("service", SERVICE);
        map.put("appid", getAppId());
        map.put("package", getContext().getPackageName());
        map.put("cpAppVersion", getAppVersionName(getContext().getPackageName()));
        map.put("model", DEVICE);
    }

    /**
     * Prepares sing-event map according to input parameters.
     *
     * @param methodName  The name of the method called
     * @param resultCode  Code of the method's result. "0" for success, others for error
     * @param currentTime Current timestamp in millisecond
     */
    private synchronized void putToSingleEventMap(final String methodName, final String resultCode,
        final long currentTime) {
        final long startTime = getOrDefault(startTimeMap, methodName, currentTime);
        final int costTime = (int) (currentTime - startTime);
        singleEventMap.put("apiName", methodName);
        singleEventMap.put("result", resultCode);
        singleEventMap.put("callTime", currentTime);
        singleEventMap.put("costTime", costTime);
        singleEventMap.put("networkType", getNetworkType());
    }

    /**
     * Prepares periodic-event map according to input parameters.
     *
     * @param methodName  The name of the method called
     * @param resultCode  Code of the method's result. "0" for success, others for error
     * @param currentTime Current timestamp in millisecond
     */
    private synchronized void putToPeriodicEventMap(final String methodName, final String resultCode,
        final long currentTime) {
        increaseResultCodeCount(methodName, resultCode);
        increaseMapValue(methodName, allCountMap);

        if (!resultCode.equals(SUCCESS)) {
            increaseMapValue(methodName, failCountMap);
        }

        final long firstReceiveTime = getOrDefault(firstReceiveTimeMap, methodName, currentTime);
        periodicEventMap.put("callTime", firstReceiveTime);

        final long lastReceiveTime = getOrDefault(lastReceiveTimeMap, methodName, currentTime);
        final int costTime = (int) (currentTime - lastReceiveTime);
        periodicEventMap.put("costTime", costTime);

        periodicEventMap.put("apiName", methodName);
        periodicEventMap.put("result", resultCodeCountMap.get(methodName));

        final long allCount = getOrDefault(allCountMap, methodName, 0L);
        periodicEventMap.put("allCnt", allCount);

        final long failCount = getOrDefault(failCountMap, methodName, 0L);
        periodicEventMap.put("failCnt", failCount);

        periodicEventMap.put("lastCallTime", currentTime);
        periodicEventMap.put("networkType", getNetworkType());

        putIfAbsent(firstReceiveTimeMap, methodName, currentTime);
        lastReceiveTimeMap.put(methodName, currentTime);
    }

    /**
     * Prepares HashMap of network type id and its human-readable string pairs.
     *
     * @return HashMap of human readable network type names
     */
    private synchronized Map<Integer, String> createNetworkTypeMap() {
        final Map<Integer, String> map = new HashMap<>();
        map.put(0, UNKNOWN);
        map.put(1, "2G");
        map.put(2, "2G");
        map.put(3, "3G");
        map.put(4, "3G");
        map.put(5, "3G");
        map.put(6, "3G");
        map.put(7, "2G");
        map.put(8, "3G");
        map.put(9, "3G");
        map.put(10, "3G");
        map.put(11, "2G");
        map.put(12, "3G");
        map.put(13, "4G");
        map.put(14, "3G");
        map.put(15, "3G");
        map.put(16, "2G");
        map.put(17, "3G");
        map.put(18, "4G");
        map.put(19, "4G");
        map.put(20, "5G");

        return Collections.unmodifiableMap(map);
    }

    /**
     * Increases count of the given result code.
     *
     * @param methodName Name of the calling method
     * @param resultCode Code of the method's result. "0" for success, others for error
     */
    private synchronized void increaseResultCodeCount(final String methodName, final String resultCode) {
        final Map<String, Long> map = getOrDefault(resultCodeCountMap, methodName, new HashMap<>());

        increaseMapValue(resultCode, map);
        resultCodeCountMap.put(methodName, map);
    }

    /**
     * Increases the value of the corresponding key which in the map.
     *
     * @param key Key for map lookup
     * @param map The Map that contains the key and its corresponding value
     */
    private synchronized void increaseMapValue(final String key, final Map<String, Long> map) {
        map.put(key, getOrDefault(map, key, 0L) + 1);
    }

    /**
     * Get the corresponding value of the key. If the key does not exist in the map then the default value is returned.
     *
     * @param map          The Map
     * @param key          Lookup key
     * @param defaultValue The default value will be returned if the key is absent
     * @param <K>          Generic type of the key
     * @param <V>          Generic type of the value
     * @return Corresponding value or default value
     */
    private synchronized <K, V> V getOrDefault(final Map<K, V> map, final K key, final V defaultValue) {
        return map.containsKey(key) ? map.get(key) : defaultValue;
    }

    /**
     * Put key-value pair to map if the key is absent.
     *
     * @param map   The Map
     * @param key   Lookup key
     * @param value The value will be put to the map if the key is absent
     * @param <K>   Generic type of the key
     * @param <V>   Generic type of the value
     */
    private synchronized <K, V> void putIfAbsent(final Map<K, V> map, final K key, final V value) {
        if (!map.containsKey(key)) {
            map.put(key, value);
        }
    }

    /**
     * Utility method that castes given object to given class type.
     *
     * @param source Source object to be casted
     * @param clazz  Class that object will be casted to its type
     * @param <S>    Source object's type
     * @param <D>    Destination type
     * @return Object that casted to D type
     */
    private synchronized <S, D> D objectCast(final S source, final Class<D> clazz) {
        return clazz.cast(source);
    }
}
`
const JAVAC_CORPACK = `package com.huawei.hms.cordova.%s.basef.handler;

import org.apache.cordova.CordovaInterface;
import org.apache.cordova.CordovaPlugin;
import org.apache.cordova.CordovaWebView;

public class CorPack {
    private final HMSLogger hmsLogger;
    private final CordovaPlugin cordovaPlugin;
    private final CordovaWebView webView;
    private final CordovaInterface cordova;
    private final CordovaEventRunner eventRunner;

    CorPack(final HMSLogger hmsLogger, final CordovaPlugin cordovaPlugin, final CordovaEventRunner eventRunner) {
        this.hmsLogger = hmsLogger;
        this.cordovaPlugin = cordovaPlugin;
        this.webView = cordovaPlugin.webView;
        this.cordova = cordovaPlugin.cordova;
        this.eventRunner = eventRunner;
    }

    public void requestPermission(int requestCode, String permission) {
        cordova.requestPermission(cordovaPlugin, requestCode, permission);
    }

    public void requestPermissions(int requestCode, String[] permissions) {
        cordova.requestPermissions(cordovaPlugin, requestCode, permissions);
    }

    public boolean hasPermission(String permission) {
        return cordova.hasPermission(permission);
    }

    public void enableLogger() {
        hmsLogger.enableLogger();
    }

    public void disableLogger() {
        hmsLogger.disableLogger();
    }

    public void startMethodExecution(String methodName) {
        hmsLogger.startMethodExecutionTimer(methodName);
    }

    public void sendSingleEvent(String methodName) {
        hmsLogger.sendSingleEvent(methodName);
    }

    public void sendSingleEvent(String methodName, String errorCode) {
        hmsLogger.sendSingleEvent(methodName, errorCode);
    }

    public CordovaWebView getCordovaWebView() {
        return webView;
    }

    public CordovaInterface getCordova() {
        return cordova;
    }

    public CordovaEventRunner getEventRunner() {
        return eventRunner;
    }
}
`
const JAVAC_CMH = `package com.huawei.hms.cordova.%s.basef.handler;

import com.huawei.hms.cordova.%s.basef.CordovaBaseModule;
import com.huawei.hms.cordova.%s.basef.CordovaEvent;
import com.huawei.hms.cordova.%s.basef.CordovaMethod;

import java.lang.reflect.Method;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class CordovaModuleHandler<T extends CordovaBaseModule> {
    private final Map<String, Method> lookupTable = new HashMap<>();
    private final List<Method> eventCache = new ArrayList<>();
    private final T instance;

    public CordovaModuleHandler(T moduleInstance) {
        this.instance = moduleInstance;
        fillLookupTable();
    }

    private void fillLookupTable(){
        Method[] methods = this.instance.getClass().getMethods();
        for(Method method : methods) {
            if (method.isAnnotationPresent(CordovaMethod.class))
                lookupTable.put(method.getName(), method);
            if(method.isAnnotationPresent(CordovaEvent.class))
                eventCache.add(method);
        }
    }

    Method getModuleMethod(String action)  {
        return lookupTable.get(action);
    }

    boolean hasModuleMethod(String action) {
        return lookupTable.containsKey(action);
    }

    List<Method> getEventCache() {
        return eventCache;
    }
    T getInstance() {
        return instance;
    }
    public Map<String, Method> getLookupTable() {
        return lookupTable;
    }

    public void clear() {
        eventCache.clear();
        lookupTable.clear();
    }
}

`
const JAVAC_CMGH = `package com.huawei.hms.cordova.%s.basef.handler;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

class CordovaModuleGroupHandler {
    private final Map<String, CordovaModuleHandler> lookupTable = new HashMap<>();
    private final List<CordovaModuleHandler> cordovaModuleHandlers;
    public CordovaModuleGroupHandler(List<CordovaModuleHandler> cordovaModuleHandlerList){
        this.cordovaModuleHandlers = cordovaModuleHandlerList;
        this.fillLookupTable();
    }

    private void fillLookupTable(){
        for(CordovaModuleHandler cordovaModuleHandler : cordovaModuleHandlers) {
            String reference = cordovaModuleHandler.getInstance().getReference();
            lookupTable.put(reference, cordovaModuleHandler);
        }
    }

    boolean hasCordovaModuleHandler(String reference) {
        return lookupTable.containsKey(reference);
    }

    CordovaModuleHandler getCordovaModuleHandler(String reference)  {
        return lookupTable.get(reference);
    }

    void clear() {
        lookupTable.clear();
        for(CordovaModuleHandler moduleHandler: cordovaModuleHandlers)
            moduleHandler.clear();
        cordovaModuleHandlers.clear();
    }
}

`
const JAVAC_COREVENTRUNNER = `package com.huawei.hms.cordova.%s.basef.handler;

import android.app.Activity;
import android.util.Log;

import org.apache.cordova.CordovaWebView;

public class CordovaEventRunner {
	private static final String TAG = CordovaEventRunner.class.getName();
	private static final String TO_STR_NOT_VALID_ERR = "Sent event parameter value is not valid! Please add toString() method to the object you " +
			"are passing or do not pass this object as an event parameter.";
	private final HMSLogger hmsLogger;
	private final CordovaWebView webView;
	private final Activity activity;

	CordovaEventRunner(final CordovaWebView cordovaWebView, final Activity activity, final HMSLogger hmsLogger) {
		this.hmsLogger = hmsLogger;
		this.webView = cordovaWebView;
		this.activity = activity;
	}

	public void sendEvent(String event, Object... params) {
		hmsLogger.sendPeriodicEvent(event);
		sendEventToJS(event, params);
	}

	public void sendEvent(String event) {
		hmsLogger.sendPeriodicEvent(event);
		sendEventToJS(event);
	}

	private void sendEventToJS(String event, Object... objects) {
		Log.i(TAG, "Periodic event " + event + " captured and event " + event + " is sending to JS.");
		StringBuilder jsFunctionBuilder = new StringBuilder();
		jsFunctionBuilder.append("javascript:");
		jsFunctionBuilder.append("window.runHMSEvent('").append(event).append("'");
		if (objects.length > 0) jsFunctionBuilder.append(buildJSEventParameters(objects));
		jsFunctionBuilder.append(");");
		activity.runOnUiThread(() -> webView.loadUrl(jsFunctionBuilder.toString()));
	}

	private String buildJSEventParameters(Object... objects) {
		StringBuilder eventParametersBuilder = new StringBuilder();
		for (Object obj : objects) {
			if (!isToStringValueValid(obj))
				Log.w(TAG, TO_STR_NOT_VALID_ERR);
			eventParametersBuilder.append(",").append(obj.toString());
		}
		return eventParametersBuilder.toString();
	}

	private boolean isToStringValueValid(Object object) {
		String originalToStr = object.getClass().getSimpleName() + "@" + Integer.toHexString(object.hashCode());
		String currentToStr = object.toString();
		return originalToStr.equals(currentToStr);
	}
}


`
const JAVAC_CORCONTROLLER = `package com.huawei.hms.cordova.%s.basef.handler;

import android.content.Intent;
import android.util.Log;

import com.huawei.hms.cordova.%s.basef.CordovaBaseModule;
import com.huawei.hms.cordova.%s.basef.HMSLog;

import org.apache.cordova.CallbackContext;
import org.apache.cordova.CordovaPlugin;
import org.json.JSONArray;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

public class CordovaController {
	private static final String TAG = CordovaController.class.getSimpleName();
	private static final String NOT_CORDOVA_MODULE_ERR = " is not a cordova module. Please check if the given module extends" +
			" class CordovaBaseModule. If the class extends CordovaBaseModule then check the main cordova class, you have to register the module" +
			" inside the main cordova class. If everything is okay so far, then please check the action parameter sent from JavaScript and ensure that" +
			" exported cordova method in java is public. If the problem persists then contact the administrator. ";

	private CordovaModuleGroupHandler groupHandler;
	private final HMSLogger hmsLogger;
	private final CordovaEventRunner eventRunner;
	private final CordovaPlugin cordovaPlugin;
	private final List<String> moduleReferences = new ArrayList<>();

	public <T extends CordovaBaseModule> CordovaController(CordovaPlugin cordovaPlugin, String service, String version, List<T> cordovaModules) {
		List<CordovaModuleHandler> moduleHandlerList = new ArrayList<>();
		for (T cordovaModule : cordovaModules) {
			CordovaModuleHandler moduleHandler = new CordovaModuleHandler(cordovaModule);
			moduleHandlerList.add(moduleHandler);
			moduleReferences.add(cordovaModule.getReference());
		}
		this.cordovaPlugin = cordovaPlugin;
		this.groupHandler = new CordovaModuleGroupHandler(moduleHandlerList);
		this.hmsLogger = HMSLogger.getInstance(cordovaPlugin.webView.getContext(), service, version);
		this.eventRunner = new CordovaEventRunner(cordovaPlugin.webView, cordovaPlugin.cordova.getActivity(), hmsLogger);

		prepareEvents();
		clearEventCache();
	}

	private void prepareEvents() {
		for (String ref : moduleReferences) {
			List eventCache = groupHandler.getCordovaModuleHandler(ref).getEventCache();
			runAllEventMethods(groupHandler.getCordovaModuleHandler(ref).getInstance(), eventCache);
		}
	}

	private <T> void runAllEventMethods(T instance, List<Method> eventCache) {
		for (Method method : eventCache) {
			try {
				method.invoke(instance, new CorPack(hmsLogger, cordovaPlugin, eventRunner));
				Log.i(TAG, "Event " + method.getName() + " is ready.");
			} catch (IllegalAccessException | InvocationTargetException e) {
				Log.e(TAG, "Event couldn't initialized. " + e.getMessage());
			}
		}
	}

	private void clearEventCache() {
		for (String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getEventCache().clear();
	}

	public boolean execute(String action, JSONArray args, final CallbackContext callbackContext) {
		if (!groupHandler.hasCordovaModuleHandler(action)) {
			Log.e(TAG, action + NOT_CORDOVA_MODULE_ERR);
			callbackContext.error("Cordova module doesn't exist.");
			return false;
		}
		CordovaModuleHandler moduleHandler = groupHandler.getCordovaModuleHandler(action);
		if (args.length() == 0 || (args.opt(0).getClass() != String.class) || !moduleHandler.hasModuleMethod(args.optString(0))) {
			Log.e(TAG, "Please ensure that the first parameter of arguments you have sent from JavaScript is the methodName and the action is the module name.");
			callbackContext.error("Function name doesn't exist.");
			return false;
		}
		String methodName = args.optString(0);
		args.remove(0); // Remove the method name after you have it.
		Method method = moduleHandler.getModuleMethod(methodName);
		Log.i(TAG, String.format(Locale.ENGLISH, "Method %s is called from module %s.", methodName, action));
		boolean isLoggerActive = false;
		if (method.isAnnotationPresent(HMSLog.class)) {
			isLoggerActive = true;
			hmsLogger.startMethodExecutionTimer(methodName);
		}
		CorPack corPack = new CorPack(hmsLogger, cordovaPlugin, eventRunner);
		Promise promise = createPromiseFromCallbackContext(callbackContext, methodName, isLoggerActive);


		try {
			method.invoke(moduleHandler.getInstance(), corPack, args, promise);
			return true;
		} catch (IllegalAccessException | IllegalArgumentException e) {
			Log.e(TAG, String.format(Locale.ENGLISH, "Error occurred when method %s in module %s was called." +
							" Exception class is %s and exception message is %s.",
					methodName, action, e.getClass().getSimpleName(), e.toString()));
			promise.error(e.toString());
		} catch (InvocationTargetException e) {
			Log.e(TAG, String.format(Locale.ENGLISH, "When method %s in module %s was called 'Invocation Target Exception' occurred." +
							" Invocation target exception means that called method was failed. Target exception is %s." +
							" Custom error message of the target exceptions is '%s.'", methodName, action, e.getTargetException().getClass(),
					e.getTargetException().getMessage()));
			promise.error(e.getTargetException().toString());
		}
		return false;
	}

	private Promise createPromiseFromCallbackContext(final CallbackContext callbackContext, String methodName, boolean isLoggerActive) {
		return new Promise(callbackContext, hmsLogger, methodName, isLoggerActive);
	}

	public void onPause(boolean multitasking) {
		Log.d(TAG, "onPause");
		for (String ref : moduleReferences) {
			groupHandler.getCordovaModuleHandler(ref).getInstance().onPause(multitasking);
		}
	}

	public void onDestroy() {
		Log.d(TAG, "onDestroy");
		for (String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onDestroy();
		groupHandler.clear();
	}

	public void onReset() {
		Log.d(TAG, "onReset");
		for (String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onReset();
	}

	public void onResume(boolean multitasking) {
		Log.d(TAG, "onResume");
		for (String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onResume(multitasking);
	}

	public void onStart() {
		Log.d(TAG, "onStart");
		for (String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onStart();
	}

	public void onStop() {
		Log.d(TAG, "onStop");
		for (String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onStop();
	}

	public void onActivityResult(int requestCode, int resultCode, Intent data) {
		Log.d(TAG, "onActivityResult");
		for(String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onActivityResult(requestCode, resultCode, data);
	}

    public void onRequestPermissionResult(int requestCode, String[] permissions, int[] grantResults) {
        Log.d(TAG, "onRequestPermissionResult");
		for(String ref : moduleReferences)
			groupHandler.getCordovaModuleHandler(ref).getInstance().onRequestPermissionResult(requestCode, permissions, grantResults);
    }
}

`

const COR_LOG = `package com.huawei.hms.cordova.%s.basef.handler;

import android.util.Log;

public class CorLog {
	private CorLog(){}
	private static boolean enable = false;

	public static void setEnable(boolean enable) {
		CorLog.enable = enable;
	}

	public static void log(int priority, String tag, String message) {
		if(enable)
			Log.println(priority, tag, message);
	}

	public static void info(String tag, String message) {
		log(Log.INFO, tag, message);
	}

	public static void warn(String tag, String message){
		log(Log.WARN, tag, message);
	}

	public static void debug(String tag, String message){
		log(Log.DEBUG, tag, message);
	}

	public static void err(String tag, String message) {
		log(Log.ERROR, tag, message);
	}

	public static void verb(String tag, String message) {
		log(Log.VERBOSE, tag, message);
	}

	public static void asst(String tag, String message){
		log(Log.ASSERT, tag, message);
	}

}

`

const PX2DP_JAVA = `package com.huawei.hms.cordova.%s.basef.util;

import android.content.res.Resources;

public class Px2Dp {
    private static final float SYSTEM_DENSITY = Resources.getSystem().getDisplayMetrics().density;

    public static int pxToDp(int px) {
        return Math.round(px * SYSTEM_DENSITY);
    }

    public static int dpToPx(int dp) {
        return Math.round(dp / SYSTEM_DENSITY);
    }
}
`
