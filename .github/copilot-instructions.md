This a tool to access Shelly Devices trough the Shelly Common HTTP API

## Shelly Common HTTP API

> Source: https://shelly-api-docs.shelly.cloud/gen1/#common-http-api

This section documents the HTTP API implemented by all Shelly devices, which defines their common traits:

-   WiFi configuration
-   Cloud settings
-   HTTP authentication settings
-   Firmware updates, identification and other system functions

For every Shelly device, one should consult this section, along with the chapter dedicated to the Shelly model in question.

## HTTP dialect

All properly formed requests return a JSON-encoded payload with an `application/json` MIME type. The meaning of values is described as **Attributes** for each documented resource.

Each resource may also accept a list of **Parameters** which should be supplied either as query-string in the URL or as `application/x-www-form-urlencoded` POST payload.

Error responses carry a 4xx HTTP response code and a `text/plain` response body, usually with an informative message for the type of error which occurred.

All resources except for `/shelly` will require Basic HTTP authentication when it is enabled via [`/settings/login`](https://shelly-api-docs.shelly.cloud/gen1/#settings-login).

The HTTP method used for performing any of the requests below is intentionally ignored. Most endpoints will always return their specific JSON payload and perform actions if query parameters are specified.

Boolean parameters can be given as `1`, `y`, `Y`, `t`, `T` or case-insensitive `true` for true, any other value will be interpreted as false.

## `/shelly`

```
<span>GET</span><span> </span><span>/shelly</span><span>

</span><span>{</span><span>
    </span><span>"type"</span><span>:</span><span> </span><span>"SHSW-21"</span><span>,</span><span>
    </span><span>"mac"</span><span>:</span><span> </span><span>"5ECF7F1632E8"</span><span>,</span><span>
    </span><span>"auth"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"fw"</span><span>:</span><span> </span><span>"20161223-111304/master@2bc16496"</span><span>,</span><span>
    </span><span>"longid"</span><span>:</span><span> </span><span>1</span><span>,</span><span>
    </span><span>"discoverable"</span><span>:</span><span> </span><span>true</span><span>
</span><span>}</span><span>
</span>
```

Provides basic information about the device. It does not require HTTP authentication, even if authentication is enabled globally. This endpoint can be used in conjunction with mDNS for device discovery and identification. It accepts no parameters.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `type` | string | Shelly model identifier |
| `mac` | string | MAC address of the device |
| `auth` | bool | Whether HTTP requests require authentication |
| `fw` | string | Current firmware version |
| `longid` | number | `1` if the device identifies itself with its full MAC address; `0` if only the last 3 bytes are used |

## `/settings`

```
<span>GET</span><span> </span><span>/settings</span><span>

</span><span>{</span><span>
    </span><span>"device"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"type"</span><span>:</span><span> </span><span>"SHSW-21"</span><span>,</span><span>
        </span><span>"mac"</span><span>:</span><span> </span><span>"16324CAABBCC"</span><span>,</span><span>
        </span><span>"hostname"</span><span>:</span><span> </span><span>"shelly1-B929CC"</span><span>
    </span><span>},</span><span>
    </span><span>"wifi_ap"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"ssid"</span><span>:</span><span> </span><span>"shellyMODEL-16324CAABBCC"</span><span>,</span><span>
        </span><span>"key"</span><span>:</span><span> </span><span>""</span><span>
    </span><span>},</span><span>
    </span><span>"wifi_sta"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
        </span><span>"ssid"</span><span>:</span><span> </span><span>"Castle"</span><span>,</span><span>
        </span><span>"ipv4_method"</span><span>:</span><span> </span><span>"dhcp"</span><span>,</span><span>
        </span><span>"ip"</span><span>:</span><span> </span><span>null</span><span>,</span><span>
        </span><span>"gw"</span><span>:</span><span> </span><span>null</span><span>,</span><span>
        </span><span>"mask"</span><span>:</span><span> </span><span>null</span><span>,</span><span>
        </span><span>"dns"</span><span>:</span><span> </span><span>null</span><span>
    </span><span>},</span><span>
    </span><span>"wifi_sta1"</span><span> </span><span>:</span><span> </span><span>{</span><span> </span><span>...</span><span> </span><span>},</span><span>
    </span><span>"ap_roaming"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"threshold"</span><span>:</span><span> </span><span>-70</span><span>
    </span><span>},</span><span>
    </span><span>"mqtt"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enable"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"server"</span><span>:</span><span> </span><span>"192.168.33.3:1883"</span><span>,</span><span>
        </span><span>"user"</span><span>:</span><span> </span><span>""</span><span>,</span><span>
        </span><span>"id"</span><span>:</span><span> </span><span>"shelly1-B929CC"</span><span>,</span><span>
        </span><span>"reconnect_timeout_max"</span><span>:</span><span> </span><span>60</span><span>,</span><span>
        </span><span>"reconnect_timeout_min"</span><span>:</span><span> </span><span>2</span><span>,</span><span>
        </span><span>"clean_session"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
        </span><span>"keep_alive"</span><span>:</span><span> </span><span>60</span><span>,</span><span>
        </span><span>"max_qos"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
        </span><span>"retain"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"update_period"</span><span>:</span><span> </span><span>30</span><span>
    </span><span>},</span><span>
    </span><span>"coiot"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
        </span><span>"update_period"</span><span>:</span><span> </span><span>15</span><span>,</span><span>
        </span><span>"peer"</span><span>:</span><span> </span><span>""</span><span>
    </span><span>},</span><span>
    </span><span>"sntp"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"server"</span><span>:</span><span> </span><span>"time.google.com"</span><span>,</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>
    </span><span>},</span><span>
    </span><span>"login"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"unprotected"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"username"</span><span>:</span><span> </span><span>"admin"</span><span>
    </span><span>},</span><span>
    </span><span>"pin_code"</span><span>:</span><span> </span><span>"123456"</span><span>,</span><span>
    </span><span>"name"</span><span>:</span><span> </span><span>"shellyMODEL-16324CAABBCC"</span><span>,</span><span>
    </span><span>"fw"</span><span>:</span><span> </span><span>"20170427-114337/master@79dbb397"</span><span>,</span><span>
    </span><span>"discoverable"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"build_info"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"build_id"</span><span>:</span><span> </span><span>"20191112-140800"</span><span>,</span><span>
        </span><span>"build_timestamp"</span><span>:</span><span> </span><span>"2019-11-12T14:08:00Z"</span><span>,</span><span>
        </span><span>"build_version"</span><span>:</span><span> </span><span>"1.0"</span><span>
    </span><span>},</span><span>
    </span><span>"cloud"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
        </span><span>"connected"</span><span>:</span><span> </span><span>true</span><span>
    </span><span>},</span><span>
    </span><span>"timezone"</span><span>:</span><span> </span><span>"Europe/Sofia"</span><span>,</span><span>
    </span><span>"lat"</span><span>:</span><span> </span><span>42.1234</span><span>,</span><span>
    </span><span>"lng"</span><span>:</span><span> </span><span>24.5678</span><span>,</span><span>
    </span><span>"tzautodetect"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"tz_utc_offset"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
    </span><span>"tz_dst"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"tz_dst_auto"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"time"</span><span>:</span><span> </span><span>"16:40"</span><span>,</span><span>
    </span><span>"unixtime"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
    </span><span>"led_status_disable"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"debug_enable"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"allow_cross_origin"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"wifirecovery_reboot_enabled"</span><span>:</span><span> </span><span>true</span><span>
</span><span>}</span><span>
</span>
```

Represents device configuration: all devices support a set of common features which are described here. Look at the device-specific `/settings` endpoint to see how each device extends it.

To configure timezone and location (for sunrise/sunset calculation) manually, set `tzautodetect` to `false`, so that custom values for `lat`, `lng` and `timezone` take effect. For a list of supported timezones, please fetch

[https://api.shelly.cloud/timezone/tzlist](https://api.shelly.cloud/timezone/tzlist)

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `device.type` | string | Device model identifier |
| `device.mac` | string | MAC address of the device in hexadecimal |
| `device.hostname` | string | Device hostname |
| `wifi_ap` | hash | WiFi access point configuration, see [`/settings/ap`](https://shelly-api-docs.shelly.cloud/gen1/#settings-ap) for details |
| `wifi_sta` | hash | WiFi client configuration. See [`/settings/sta`](https://shelly-api-docs.shelly.cloud/gen1/#settings-sta) for details |
| `wifi_sta1` | hash | Alternative WiFi client configuration. See [`/settings/sta`](https://shelly-api-docs.shelly.cloud/gen1/#settings-sta) for details |
| `ap_roaming.enabled` | bool | AP roaming enabled flag |
| `ap_roaming.threshold` | number | AP roaming threshold, `dBm` |
| `mqtt` | hash | Contains MQTT-related settings |
| `coiot.enabled` | bool | CoIoT enabled flag |
| `coiot.update_period` | number | Update period of CoIoT messages, `s` |
| `coiot.peer` | string | CoIoT peer (in format `ip:port`, empty means `mcast`) |
| `sntp.server` | string | Time server host |
| `sntp.enabled` | bool | SNTP enabled flag |
| `login` | hash | Credentials used for HTTP Basic authentication for the REST interface. If `login.enabled` is `true` clients must include an `Authorization: Basic ...` HTTP header with valid credentials when performing TP requests |
| `pin_code` | string | Current generated PIN code |
| `name` | string | Unique name of the device |
| `fw` | string | Current FW version |
| `discoverable` | bool | Device discoverable (i.e. visible) flag |
| `build_info` | hash | Build info |
| `cloud.enabled` | bool | Cloud enabled flag |
| `cloud.connected` | bool | Cloud connected flag |
| `timezone` | string | Timezone identifier |
| `lat` | number | Degrees latitude in decimal format, South is negative |
| `lng` | number | Degrees longitude in decimal fomrat, between -180° and 180° |
| `tzautodetect` | bool | Timezone auto-detect enabled |
| `tz_utc_offset` | number | UTC offset |
| `tz_dst` | bool | Daylight saving time |
| `tz_dst_auto` | bool | Auto update daylight saving time |
| `time` | string | Current time in HH:MM format if synced |
| `unixtime` | number | Unix timestamp if synced; `0` otherwise |
| `led_status_disable` | bool | **For Dimmer 1/2, DW2, i3, RGBW2, Plug, Plug S, EM, 3EM, 1L, 1PM, 2.5 and Button1** Whether LED indication for network status is enabled |
| `fw_mode` | string | Current firmware mode, only for devices where this is changeable (e.g. Shelly RGBW2 Color/RGBW2 White) |
| `debug_enable` | bool | Debug file logger enabled flag |
| `allow_cross_origin` | bool | HTTP Cross-Origin Resource Sharing allowed flag |
| `wifirecovery_reboot_enabled` | bool | Whether WiFi-Recovery reboot is enabled. Only applicable for Shelly 1/1PM, Shelly 2, Shelly 2.5, Shelly Plug/PlugS, Shelly EM, Shelly 3EM |

### Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `reset` | bool | Will perform a factory reset of the device |
| `ap_roaming_enabled` | bool | Enable/disable AP roaming |
| `ap_roaming_threshold` | number | Set AP roaming threshold, `dBm` |
| `mqtt_enable` | bool | Enable connecting to a MQTT broker |
| `mqtt_server` | string | MQTT broker IP address and port, ex. `10.0.0.1:1883` |
| `mqtt_clean_session` | bool | MQTT clean session flag |
| `mqtt_retain` | bool | MQTT retain flag |
| `mqtt_user` | string | MQTT username, leave empty to disable authentication |
| `mqtt_pass` | string | MQTT password |
| `mqtt_id` | string | MQTT ID -- by default this has the form `<shellymodel>-<deviceid>` e.g. `shelly1-B929CC`. If you wish to use custom a MQTT ID, it is recommended that it doesn't exceed 25 characters. |
| `mqtt_reconnect_timeout_max` | number | Maximum interval for reconnect attempts |
| `mqtt_reconnect_timeout_min` | number | Minimum interval for reconnect attempts |
| `mqtt_keep_alive` | number | MQTT keep alive period in seconds |
| `mqtt_update_period` | number | Periodic update in seconds, `0` to disable |
| `mqtt_max_qos` | number | Max value of QOS for MQTT packets |
| `coiot_enable` | bool | Enable/disable CoIoT |
| `coiot_update_period` | number | Update period of CoIoT messages `15..65535s` |
| `coiot_peer` | string | Set to either `mcast` to switch CoIoT to multicast or to `ip[:port]` to switch CoIoT to unicast (`port` is optional, default is 5683) |
| `sntp_server` | string | Time server host to be used instead of the default `time.google.com`. An empty value disables timekeeping and requires reboot to apply. |
| `name` | string | User-configurable device name |
| `discoverable` | bool | Set whether device should be discoverable (i.e. visible) |
| `timezone` | string | Timezone identifier |
| `lat` | number | Degrees latitude in decimal format, South is negative |
| `lng` | number | Degrees longitude in decimal format, `-180°..180°` |
| `tzautodetect` | bool | Set this to `false` if you want to set custom geolocation (`lat` and `lng`) or custom `timezone`. |
| `tz_utc_offset` | number | UTC offset |
| `tz_dst` | number | Daylight saving time `0/1` |
| `tz_dst_auto` | number | Auto update daylight saving time `0/1` |
| `led_status_disable` | bool | **For Dimmer 1/2, DW2, i3, RGBW2, Plug, Plug S, EM, 3EM, 1L, 1PM, 2.5 and Button1** Enable/Disable LED indication for network status |
| `debug_enable` | bool | Enable/disable debug file logger |
| `allow_cross_origin` | bool | Allow/forbid HTTP Cross-Origin Resource Sharing |
| `wifirecovery_reboot_enabled` | bool | Enable/disbale WiFi-Recovery reboot. Only applicable for Shelly 1/1PM, Shelly 2, Shelly 2.5, Shelly Plug/PlugS, Shelly EM, Shelly 3EM |

## `/settings/ap`

```
<span>GET</span><span> </span><span>/settings/ap</span><span>

</span><span>{</span><span>
    </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"ssid"</span><span>:</span><span> </span><span>"shellyswitch-163248"</span><span>,</span><span>
    </span><span>"key"</span><span>:</span><span> </span><span>""</span><span>
</span><span>}</span><span>
</span>
```

Provides information about the current WiFi AP configuration and allows changes. The returned document is identical to the one returned by `/settings` in the `wifi_ap` key. Shelly devices do not allow the SSID for AP WiFi mode to be changed.

Parameters are applied immediately. Setting the `enabled` flag for AP mode to `1` will automatically disable STA mode.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `enabled` | bool | Whether AP mode is active |
| `ssid` | string | SSID created by the device's AP |
| `key` | string | WiFi password required for association with the device's AP |

### Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `enabled` | bool | Set to `1` to return the device to AP WiFi mode |
| `key` | string | WiFi password required for association with the device's AP |

## `/settings/sta`

```
<span>GET</span><span> </span><span>/settings/sta</span><span>

</span><span>{</span><span>
    </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"ssid"</span><span>:</span><span> </span><span>"Castle"</span><span>,</span><span>
    </span><span>"ipv4_method"</span><span>:</span><span> </span><span>"dhcp"</span><span>,</span><span>
    </span><span>"ip"</span><span>:</span><span> </span><span>null</span><span>,</span><span>
    </span><span>"gw"</span><span>:</span><span> </span><span>null</span><span>,</span><span>
    </span><span>"mask"</span><span>:</span><span> </span><span>null</span><span>,</span><span>
    </span><span>"dns"</span><span>:</span><span> </span><span>null</span><span>
</span><span>}</span><span>
</span>
```

Provides information about the current WiFi Client mode configuration and allows changes. The returned document is identical to the one returned by `/settings` in the `wifi_sta` key.

Parameters are applied immediately. Setting the `enabled` flag for STA mode to `1` will automatically disable AP mode.

An identical resource is available at `/settings/sta1`. This allows for devices to have a second WiFi STA configuration for fallback, and will cycle between the two when the one currently selected becomes unavailable.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `enabled` | bool | Whether STA mode is active |
| `ssid` | string | SSID of STA the device will associate with |
| `ipv4_method` | string | `dhcp` or `static` |
| `ip` | string | Local IP address if `ipv4_method` is `static` |
| `gw` | string | Local gateway IP address if `ipv4_method` is `static` |
| `mask` | string | Mask if `ipv4_method` is `static` |
| `dns` | string | DNS address if `ipv4_method` is `static` |

### Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `enabled` | bool | Set to `1` to make STA the current WiFi mode |
| `ssid` | string | The WiFi SSID to associate with |
| `key` | string | The password required for associating to the given WiFi SSID |
| `ipv4_method` | string | `dhcp` or `static` |
| `ip` | string | Local IP address if `ipv4_method` is `static` |
| `netmask` | string | Mask if `ipv4_method` is `static` |
| `gateway` | string | Local gateway IP address if `ipv4_method` is `static` |
| `dns` | string | DNS address if `ipv4_method` is `static` |

Sinve v.1.7.0, setting a static IP only requires `netmask` to be specified, `gateway` can be left empty.

Please note that cloud must be disabled manually before setting a static IP config without a `gateway` and enabled manually when switching back to `dhcp` config ([`/settings/cloud`](https://shelly-api-docs.shelly.cloud/gen1/#settings-cloud)).

## `/settings/login`

```
<span>GET</span><span> </span><span>/settings/login</span><span>

</span><span>{</span><span>
    </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"unprotected"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"username"</span><span>:</span><span> </span><span>"admin"</span><span>
</span><span>}</span><span>
</span>
```

```
<span>GET</span><span> </span><span>/settings/login?enabled=</span><span>1</span><span>&amp;username=boss&amp;password=thebigone</span><span>

</span><span>{</span><span>
    </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"unprotected"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"username"</span><span>:</span><span> </span><span>"boss"</span><span>
</span><span>}</span><span>
</span>
```

HTTP authentication configuration: `enabled` flag, credentials. `unprotected` is initially false and is used by the user interface to show a warning when auth is disabled. If the user wants to keep using Shelly without a password, they can set `unprotected` to hide the warning.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `enabled` | bool | Whether HTTP authentication is required |
| `unprotected` | bool | Whether the user is aware of the risks |
| `username` | string | Username |

In order to prevent security issues (e.g. when forwarding logs to 3rd parties), since v1.7.0 the `password` attribute for login protected devices is no longer returned on the API call.

### Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `enabled` | bool | Whether to require HTTP authentication |
| `unprotected` | bool | Whether the user is aware of the risks |
| `username` | string | Length between 1 and 50 |
| `password` | string | Length between 1 and 50 |

## `/settings/cloud`

```
<span>GET</span><span> </span><span>/settings/cloud?enabled=</span><span>1</span><span>

</span><span>{</span><span>
    </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>
</span><span>}</span><span>
</span>
```

Can set the `enabled` (connect to cloud) flag. When set, Shelly will keep a secure connection to Allterco's servers and allow monitoring and control from anywhere.

## `/settings/actions`

```
<span>GET</span><span> </span><span>/settings/actions</span><span>

</span><span>{</span><span>
  </span><span>"actions"</span><span>:</span><span> </span><span>{</span><span>
    </span><span>"out_on_url"</span><span>:</span><span> </span><span>[</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[</span><span>
          </span><span>"http://192.168.1.4/on"</span><span>,</span><span>
          </span><span>"http://192.168.1.5/on"</span><span>
        </span><span>],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>true</span><span>
      </span><span>}</span><span>
    </span><span>],</span><span>
    </span><span>"out_off_url"</span><span>:</span><span> </span><span>[</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>
      </span><span>}</span><span>
    </span><span>],</span><span>
    </span><span>"over_power_url"</span><span>:</span><span> </span><span>[</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"over_power_url_threshold"</span><span>:</span><span> </span><span>0</span><span>
      </span><span>},</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>1</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"over_power_url_threshold"</span><span>:</span><span> </span><span>0</span><span>
      </span><span>},</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>2</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"over_power_url_threshold"</span><span>:</span><span> </span><span>0</span><span>
      </span><span>}</span><span>
    </span><span>],</span><span>
    </span><span>"under_power_url"</span><span>:</span><span> </span><span>[</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"under_power_url_threshold"</span><span>:</span><span> </span><span>0</span><span>
      </span><span>},</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>1</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"under_power_url_threshold"</span><span>:</span><span> </span><span>0</span><span>
      </span><span>},</span><span>
      </span><span>{</span><span>
        </span><span>"index"</span><span>:</span><span> </span><span>2</span><span>,</span><span>
        </span><span>"urls"</span><span>:</span><span> </span><span>[],</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"under_power_url_threshold"</span><span>:</span><span> </span><span>0</span><span>
      </span><span>}</span><span>
    </span><span>]</span><span>
  </span><span>}</span><span>
</span><span>}</span><span>
</span>
```

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `index` | number | Channel number |
| `name` | string | Action name |
| `enabled` | bool | Enable/disable action |
| `urls[]` | string | Action URL |

### Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `index` | number | Channel number |
| `name` | string | Action name |
| `enabled` | bool | Enable/disable action |
| `urls[]` | string | Action URL |

Since v1.8.0 devices support multiple actions. Max number of URLs per action is 5. Max URLs length is 256.

### Examples:

-   Setting two URLs for `out_on_url` action on channel 0:

`http://192.168.33.1/settings/actions?index=0&name=out_on_url&enabled=true&urls[]=http://192.168.1.4/on&urls[]=http://192.168.1.5/on`

-   Disable only action `out_on_url` on channel 0:

`http://192.168.33.1/settings/actions?index=0&name=out_on_url&enabled=false`

-   Delete URLs for action `out_on_url` action on channel 0:

`http://192.168.33.1/settings/actions?index=0&name=out_on_url&urls[]=`

-   Some devices have additional parameters which may be set along with actions. Example: Setting URLs and `over_power_url_threshold` parameter for action `over_power_url` on channel 0:

`http://192.168.33.1/settings/actions?index=0&name=over_power_url&enabled=true&urls[]=192.168.1.4/overpower&urls[]=192.168.1.5/overpower&over_power_url_threshold=100`

## `/status`

```
<span>GET</span><span> </span><span>/status</span><span>

</span><span>{</span><span>
    </span><span>"wifi_sta"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"connected"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
        </span><span>"ssid"</span><span>:</span><span> </span><span>"Castle"</span><span>,</span><span>
        </span><span>"ip"</span><span>:</span><span> </span><span>"192.168.2.65"</span><span>,</span><span>
        </span><span>"rssi"</span><span>:</span><span> </span><span>-54</span><span>
    </span><span>},</span><span>
    </span><span>"cloud"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"enabled"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
        </span><span>"connected"</span><span>:</span><span> </span><span>false</span><span>
    </span><span>},</span><span>
    </span><span>"mqtt"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"connected"</span><span>:</span><span> </span><span>false</span><span>
    </span><span>},</span><span>
    </span><span>"time"</span><span>:</span><span> </span><span>"17:42"</span><span>,</span><span>
    </span><span>"unixtime"</span><span>:</span><span> </span><span>0</span><span>,</span><span>
    </span><span>"serial"</span><span>:</span><span> </span><span>1</span><span>,</span><span>
    </span><span>"has_update"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
    </span><span>"mac"</span><span>:</span><span> </span><span>"98F4ABB929CC"</span><span>,</span><span>
    </span><span>"update"</span><span>:</span><span> </span><span>{</span><span>
        </span><span>"status"</span><span>:</span><span> </span><span>"pending"</span><span>,</span><span>
        </span><span>"has_update"</span><span>:</span><span> </span><span>true</span><span>,</span><span>
        </span><span>"new_version"</span><span>:</span><span> </span><span>"20200320-123430/v1.6.2@514044b4"</span><span>,</span><span>
        </span><span>"old_version"</span><span>:</span><span> </span><span>"20200320-123430/v1.6.2@514044b4"</span><span>
    </span><span>},</span><span>
    </span><span>"ram_total"</span><span>:</span><span> </span><span>50648</span><span>,</span><span>
    </span><span>"ram_free"</span><span>:</span><span> </span><span>38376</span><span>,</span><span>
    </span><span>"ram_lwm"</span><span>:</span><span> </span><span>32968</span><span>,</span><span>
    </span><span>"fs_size"</span><span>:</span><span> </span><span>233681</span><span>,</span><span>
    </span><span>"fs_free"</span><span>:</span><span> </span><span>174194</span><span>,</span><span>
    </span><span>"uptime"</span><span>:</span><span> </span><span>39</span><span>
</span><span>}</span><span>
</span>
```

Encapsulates current device status information. While settings can generally be modified and don't react to the environment, this endpoint provides information about transient data which may change due to external conditions.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `wifi_sta` | hash | Current status of the WiFi connection |
| `wifi_sta.ip` | string | IP address assigned to this device by the WiFi router |
| `cloud` | hash | Current cloud connection status |
| `mqtt.connected` | bool | MQTT connection status, when MQTT is enabled |
| `time` | string | The current hour and minutes, in `HH:MM` format |
| `unixtime` | number | Unix timestamp if synced; `0` otherwise |
| `serial` | number | Cloud serial number |
| `update` | hash | Info whether newer firmware version is available, see [`/ota`](https://shelly-api-docs.shelly.cloud/gen1/#ota) for more details |
| `ram_total` | number | Total amount of system memory in bytes |
| `ram_free` | number | Available amount of system memory in bytes |
| `ram_lwm` | number | Minimal watermark of the system free memory in bytes |
| `fs_size` | number | Total amount of the file system in bytes |
| `fs_free` | number | Available amount of the file system in bytes |
| `uptime` | number | Seconds elapsed since boot |

## `/reboot`

When requested will cause a reboot of the device.

## `/ota`

```
<span>GET</span><span> </span><span>/ota</span><span>

</span><span>{</span><span>
    </span><span>"status"</span><span>:</span><span> </span><span>"idle"</span><span>,</span><span>
    </span><span>"has_update"</span><span>:</span><span> </span><span>false</span><span>,</span><span>
    </span><span>"new_version"</span><span>:</span><span> </span><span>"20200320-123430/v1.6.2@514044b4"</span><span>,</span><span>
    </span><span>"old_version"</span><span>:</span><span> </span><span>"20200320-123430/v1.6.2@514044b4"</span><span>
</span><span>}</span><span>
</span>
```

Provides information about the device firmware version and the ability to trigger an over-the-air update.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `status` | string | One of `idle`, `pending`, `updating`, `unknown` |
| `has_update` | bool | Whether an update is available |
| `new_version` | string | ID of new firmware version |
| `old_version` | string | ID of old (current) firmware version |
| `beta_version` | string | ID of beta firmware if one is available. This attribute is optional. |

### Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| `url` | string | Run firmware update from specified URL |
| `update` | bool | Run firmware update from default URL |
| `beta` | bool | Run firmware update from beta URL (if available) |

## `/ota/check`

```
<span>GET</span><span> </span><span>/ota/check</span><span>

</span><span>{</span><span>
    </span><span>"status"</span><span>:</span><span> </span><span>"ok"</span><span>
</span><span>}</span><span>
</span>
```

Manual check for new firmware version.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `status` | string | `ok` or `running` |

## `/wifiscan`

```
<span>GET</span><span> </span><span>/wifiscan</span><span>

</span><span>{</span><span>
    </span><span>"wifiscan"</span><span>:</span><span> </span><span>"done"</span><span>,</span><span>
    </span><span>"results"</span><span>:</span><span> </span><span>[</span><span>
        </span><span>{</span><span>
            </span><span>"ssid"</span><span>:</span><span> </span><span>"some_network_1"</span><span>,</span><span>
            </span><span>"auth"</span><span>:</span><span> </span><span>4</span><span>,</span><span>
            </span><span>"channel"</span><span>:</span><span> </span><span>1</span><span>,</span><span>
            </span><span>"bssid"</span><span>:</span><span> </span><span>"XXXXXXXXXXXX"</span><span>,</span><span>
            </span><span>"rssi"</span><span>:</span><span> </span><span>-70</span><span>
        </span><span>},</span><span>
        </span><span>{</span><span>
            </span><span>"ssid"</span><span>:</span><span> </span><span>"some_network_2"</span><span>,</span><span>
            </span><span>"auth"</span><span>:</span><span> </span><span>3</span><span>,</span><span>
            </span><span>"channel"</span><span>:</span><span> </span><span>6</span><span>,</span><span>
            </span><span>"bssid"</span><span>:</span><span> </span><span>"YYYYYYYYYYYY"</span><span>,</span><span>
            </span><span>"rssi"</span><span>:</span><span> </span><span>-83</span><span>
        </span><span>}</span><span>
    </span><span>]</span><span>
</span><span>}</span><span>
</span>
```

Only available when in AP mode. Starts a WiFi scan and provides information about found networks.

### Attributes

| Attribute | Type | Description |
| --- | --- | --- |
| `wifiscan` | string | One of `failed`, `done`, `not AP mode`, `started`, `inprogress` |
| `results.{index}.ssid` | string | SSID |
| `results.{index}.auth` | number | Auth mode: `0`\=open, `1`\=WEP, `2`\=WPA PSK, `3`\=WPA2 PSK, `4`\=WPA WPA2 PSK, `5`\=WPA2 ENTERPRISE |
| `results.{index}.channel` | number | Channel |
| `results.{index}.bssid` | string | BSSID in hex |
| `results.{index}.rssi` | number | Signal strength |

## `/debug/log`

Retrieve debug log file 0

## `/debug/log1`

Retrieve debug log file 1

## `/cit/d`

```
<span>GET</span><span> </span><span>/cit/d</span><span>

</span><span>{</span><span>
    </span><span>"blk"</span><span>:[</span><span>
        </span><span>...</span><span>
    </span><span>],</span><span>
    </span><span>"sen"</span><span>:[</span><span>
        </span><span>...</span><span>
    </span><span>]</span><span>
</span><span>}</span><span>
</span>
```

Get CoIoT description package over HTTP

## `/sta_cache_reset`

Reset STA cache: only applicable on devices which support STA caching (battery operated devices Shelly Door/Window 1/2, Shelly H&T, Shelly Smoke, Shelly Flood, Shelly Button1)

## MQTT Support

Shelly devices include basic MQTT support since version 1.3.0. While many device settings are only available over HTTP, MQTT allows for real-time monitoring and eases integration with external systems.

## MQTT Configuration

To configure a Shelly device for MQTT, set the connection parameters via the Shelly App, web interface or HTTP [`/settings`](https://shelly-api-docs.shelly.cloud/gen1/#settings) endpoint. Note, that enabling MQTT will disable Allterco's cloud service. Shelly devices do not support secure MQTT connections.

Mandatory settings are:

-   `mqtt_server` is your broker's `address:port`
-   `mqtt_enable` needs to be set to `true`

In case authentication is required, `mqtt_user` and `mqtt_pass` must also be set. In certain scenarios, it may be desirable to set `mqtt_max_qos` and `mqtt_retain` to prevent loss of data.

By default, the device's MQTT ID is `<shellymodel>-<deviceid>`, for example `shelly1-B929CC`. The MQTT ID can be changed via the `mqtt_id` parameter in [`/settings`](https://shelly-api-docs.shelly.cloud/gen1/#settings). If you wish to use custom a MQTT ID, it is recommended that it doesn't exceed 25 characters.

In order to apply the MQTT configuration, the device requires a reboot.

## Availability and announces

Since v1.4.3: on MQTT connect, Shellies will publish:

-   an announce message on `shellies/announce` and, since v1.6.0, on `shellies/<shellymodel>-<deviceid>/announce`. The message is JSON-formatted and contains a list of attributes: `id`, `mode` (if applicable), `model`, `mac`, `ip`, `new_fw` (`true` when an update is available), `fw_ver` (contains the current firmware version);
-   availability message on `shellies/<shellymodel>-<deviceid>/online` with payload `true`;
-   complete current state. This is device specific and is described in detail for each device below.

Device state is reported periodically, every 30 seconds by default. This can be changed by setting a new period for updates: `mqtt_update_period` under [`/settings`](https://shelly-api-docs.shelly.cloud/gen1/#settings). A value of `0` will disable periodic updates.

Default LWT topic and message are `shellies/<shellymodel>-<deviceid>/online`, `false`. If these are not set after a firmware upgrade -- perform a factory reset of the device. The LWT topic is retained on user configuration (if the Retain flag is set). However, we do not recommend using retained MQTT messages.

## Common MQTT commands

Shellies support a set of commands published on `shellies/command` or `shellies/<shellymodel>-<deviceid>/command` to address an individual device:

-   `announce` will trigger:
    -   an announce packet by every Shelly connected to the broker on `shellies/announce` and, since v1.6.0, on `shellies/<shellymodel>-<deviceid>/announce`;
    -   since v1.8.0, `shellies/<shellymodel>-<deviceid>/info` with the content of the http `/status` endpoint.
-   `update` will cause all Shellies to publish their state;
-   `update_fw` to perform a firmware update when one is available.

Each Shelly model exports it's own set of topics for monitoring and control, all structured under `/shellies/<shellymodel>-<deviceid>`. Please see the respective product docs for details on the protocol:

-   [Shelly1/1PM](https://shelly-api-docs.shelly.cloud/gen1/#shelly1-1pm-mqtt)
-   [Shelly2](https://shelly-api-docs.shelly.cloud/gen1/#shelly2-mqtt)
-   [Shelly2.5](https://shelly-api-docs.shelly.cloud/gen1/#shelly2-5-mqtt)
-   [Shelly4Pro](https://shelly-api-docs.shelly.cloud/gen1/#shelly4pro-mqtt)
-   [Shelly Plug/Plug S](https://shelly-api-docs.shelly.cloud/gen1/#shelly-plug-plugs-mqtt)
-   [Shelly i3](https://shelly-api-docs.shelly.cloud/gen1/#shelly-i3-mqtt)
-   [Shelly Button1](https://shelly-api-docs.shelly.cloud/gen1/#shelly-button1-mqtt)
-   [Shelly Bulb](https://shelly-api-docs.shelly.cloud/gen1/#shelly-bulb-mqtt)
-   [Shelly Vintage](https://shelly-api-docs.shelly.cloud/gen1/#shelly-vintage-mqtt)
-   [Shelly Duo](https://shelly-api-docs.shelly.cloud/gen1/#shelly-duo-mqtt)
-   [Shelly RGBW2 Color](https://shelly-api-docs.shelly.cloud/gen1/#shelly-rgbw2-color-mqtt)
-   [Shelly RGBW2 White](https://shelly-api-docs.shelly.cloud/gen1/#shelly-rgbw2-white-mqtt)
-   [Shelly Dimmer/SL](https://shelly-api-docs.shelly.cloud/gen1/#shelly-dimmer-sl-mqtt)
-   [Shelly Sense](https://shelly-api-docs.shelly.cloud/gen1/#shelly-sense-mqtt)
-   [Shelly H&T](https://shelly-api-docs.shelly.cloud/gen1/#shelly-h-amp-t-mqtt)
-   [Shelly Smoke](https://shelly-api-docs.shelly.cloud/gen1/#shelly-smoke-mqtt)
-   [Shelly Flood](https://shelly-api-docs.shelly.cloud/gen1/#shelly-flood-mqtt)
-   [Shelly Door/Window](https://shelly-api-docs.shelly.cloud/gen1/#shelly-door-window-mqtt)
-   [Shelly Gas](https://shelly-api-docs.shelly.cloud/gen1/#shelly-gas-mqtt)
-   [Shelly EM](https://shelly-api-docs.shelly.cloud/gen1/#shelly-em-mqtt)
-   [Shelly 3EM](https://shelly-api-docs.shelly.cloud/gen1/#shelly-3em-mqtt)

