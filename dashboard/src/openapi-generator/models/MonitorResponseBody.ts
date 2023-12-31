/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { EnduroMonitorPingEventResponseBody } from './EnduroMonitorPingEventResponseBody';
import {
    EnduroMonitorPingEventResponseBodyFromJSON,
    EnduroMonitorPingEventResponseBodyFromJSONTyped,
    EnduroMonitorPingEventResponseBodyToJSON,
} from './EnduroMonitorPingEventResponseBody';
import type { EnduroPackageCreatedEventResponseBody } from './EnduroPackageCreatedEventResponseBody';
import {
    EnduroPackageCreatedEventResponseBodyFromJSON,
    EnduroPackageCreatedEventResponseBodyFromJSONTyped,
    EnduroPackageCreatedEventResponseBodyToJSON,
} from './EnduroPackageCreatedEventResponseBody';
import type { EnduroPackageLocationUpdatedEventResponseBody } from './EnduroPackageLocationUpdatedEventResponseBody';
import {
    EnduroPackageLocationUpdatedEventResponseBodyFromJSON,
    EnduroPackageLocationUpdatedEventResponseBodyFromJSONTyped,
    EnduroPackageLocationUpdatedEventResponseBodyToJSON,
} from './EnduroPackageLocationUpdatedEventResponseBody';
import type { EnduroPackageStatusUpdatedEventResponseBody } from './EnduroPackageStatusUpdatedEventResponseBody';
import {
    EnduroPackageStatusUpdatedEventResponseBodyFromJSON,
    EnduroPackageStatusUpdatedEventResponseBodyFromJSONTyped,
    EnduroPackageStatusUpdatedEventResponseBodyToJSON,
} from './EnduroPackageStatusUpdatedEventResponseBody';
import type { EnduroPackageUpdatedEventResponseBody } from './EnduroPackageUpdatedEventResponseBody';
import {
    EnduroPackageUpdatedEventResponseBodyFromJSON,
    EnduroPackageUpdatedEventResponseBodyFromJSONTyped,
    EnduroPackageUpdatedEventResponseBodyToJSON,
} from './EnduroPackageUpdatedEventResponseBody';
import type { EnduroPreservationActionCreatedEventResponseBody } from './EnduroPreservationActionCreatedEventResponseBody';
import {
    EnduroPreservationActionCreatedEventResponseBodyFromJSON,
    EnduroPreservationActionCreatedEventResponseBodyFromJSONTyped,
    EnduroPreservationActionCreatedEventResponseBodyToJSON,
} from './EnduroPreservationActionCreatedEventResponseBody';
import type { EnduroPreservationActionUpdatedEventResponseBody } from './EnduroPreservationActionUpdatedEventResponseBody';
import {
    EnduroPreservationActionUpdatedEventResponseBodyFromJSON,
    EnduroPreservationActionUpdatedEventResponseBodyFromJSONTyped,
    EnduroPreservationActionUpdatedEventResponseBodyToJSON,
} from './EnduroPreservationActionUpdatedEventResponseBody';
import type { EnduroPreservationTaskCreatedEventResponseBody } from './EnduroPreservationTaskCreatedEventResponseBody';
import {
    EnduroPreservationTaskCreatedEventResponseBodyFromJSON,
    EnduroPreservationTaskCreatedEventResponseBodyFromJSONTyped,
    EnduroPreservationTaskCreatedEventResponseBodyToJSON,
} from './EnduroPreservationTaskCreatedEventResponseBody';
import type { EnduroPreservationTaskUpdatedEventResponseBody } from './EnduroPreservationTaskUpdatedEventResponseBody';
import {
    EnduroPreservationTaskUpdatedEventResponseBodyFromJSON,
    EnduroPreservationTaskUpdatedEventResponseBodyFromJSONTyped,
    EnduroPreservationTaskUpdatedEventResponseBodyToJSON,
} from './EnduroPreservationTaskUpdatedEventResponseBody';

/**
 * MonitorResponseBody result type (default view)
 * @export
 * @interface MonitorResponseBody
 */
export interface MonitorResponseBody {
    /**
     * 
     * @type {EnduroMonitorPingEventResponseBody}
     * @memberof MonitorResponseBody
     */
    monitorPingEvent?: EnduroMonitorPingEventResponseBody;
    /**
     * 
     * @type {EnduroPackageCreatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    packageCreatedEvent?: EnduroPackageCreatedEventResponseBody;
    /**
     * 
     * @type {EnduroPackageLocationUpdatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    packageLocationUpdatedEvent?: EnduroPackageLocationUpdatedEventResponseBody;
    /**
     * 
     * @type {EnduroPackageStatusUpdatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    packageStatusUpdatedEvent?: EnduroPackageStatusUpdatedEventResponseBody;
    /**
     * 
     * @type {EnduroPackageUpdatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    packageUpdatedEvent?: EnduroPackageUpdatedEventResponseBody;
    /**
     * 
     * @type {EnduroPreservationActionCreatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    preservationActionCreatedEvent?: EnduroPreservationActionCreatedEventResponseBody;
    /**
     * 
     * @type {EnduroPreservationActionUpdatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    preservationActionUpdatedEvent?: EnduroPreservationActionUpdatedEventResponseBody;
    /**
     * 
     * @type {EnduroPreservationTaskCreatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    preservationTaskCreatedEvent?: EnduroPreservationTaskCreatedEventResponseBody;
    /**
     * 
     * @type {EnduroPreservationTaskUpdatedEventResponseBody}
     * @memberof MonitorResponseBody
     */
    preservationTaskUpdatedEvent?: EnduroPreservationTaskUpdatedEventResponseBody;
}

/**
 * Check if a given object implements the MonitorResponseBody interface.
 */
export function instanceOfMonitorResponseBody(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MonitorResponseBodyFromJSON(json: any): MonitorResponseBody {
    return MonitorResponseBodyFromJSONTyped(json, false);
}

export function MonitorResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): MonitorResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'monitorPingEvent': !exists(json, 'monitor_ping_event') ? undefined : EnduroMonitorPingEventResponseBodyFromJSON(json['monitor_ping_event']),
        'packageCreatedEvent': !exists(json, 'package_created_event') ? undefined : EnduroPackageCreatedEventResponseBodyFromJSON(json['package_created_event']),
        'packageLocationUpdatedEvent': !exists(json, 'package_location_updated_event') ? undefined : EnduroPackageLocationUpdatedEventResponseBodyFromJSON(json['package_location_updated_event']),
        'packageStatusUpdatedEvent': !exists(json, 'package_status_updated_event') ? undefined : EnduroPackageStatusUpdatedEventResponseBodyFromJSON(json['package_status_updated_event']),
        'packageUpdatedEvent': !exists(json, 'package_updated_event') ? undefined : EnduroPackageUpdatedEventResponseBodyFromJSON(json['package_updated_event']),
        'preservationActionCreatedEvent': !exists(json, 'preservation_action_created_event') ? undefined : EnduroPreservationActionCreatedEventResponseBodyFromJSON(json['preservation_action_created_event']),
        'preservationActionUpdatedEvent': !exists(json, 'preservation_action_updated_event') ? undefined : EnduroPreservationActionUpdatedEventResponseBodyFromJSON(json['preservation_action_updated_event']),
        'preservationTaskCreatedEvent': !exists(json, 'preservation_task_created_event') ? undefined : EnduroPreservationTaskCreatedEventResponseBodyFromJSON(json['preservation_task_created_event']),
        'preservationTaskUpdatedEvent': !exists(json, 'preservation_task_updated_event') ? undefined : EnduroPreservationTaskUpdatedEventResponseBodyFromJSON(json['preservation_task_updated_event']),
    };
}

export function MonitorResponseBodyToJSON(value?: MonitorResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'monitor_ping_event': EnduroMonitorPingEventResponseBodyToJSON(value.monitorPingEvent),
        'package_created_event': EnduroPackageCreatedEventResponseBodyToJSON(value.packageCreatedEvent),
        'package_location_updated_event': EnduroPackageLocationUpdatedEventResponseBodyToJSON(value.packageLocationUpdatedEvent),
        'package_status_updated_event': EnduroPackageStatusUpdatedEventResponseBodyToJSON(value.packageStatusUpdatedEvent),
        'package_updated_event': EnduroPackageUpdatedEventResponseBodyToJSON(value.packageUpdatedEvent),
        'preservation_action_created_event': EnduroPreservationActionCreatedEventResponseBodyToJSON(value.preservationActionCreatedEvent),
        'preservation_action_updated_event': EnduroPreservationActionUpdatedEventResponseBodyToJSON(value.preservationActionUpdatedEvent),
        'preservation_task_created_event': EnduroPreservationTaskCreatedEventResponseBodyToJSON(value.preservationTaskCreatedEvent),
        'preservation_task_updated_event': EnduroPreservationTaskUpdatedEventResponseBodyToJSON(value.preservationTaskUpdatedEvent),
    };
}

