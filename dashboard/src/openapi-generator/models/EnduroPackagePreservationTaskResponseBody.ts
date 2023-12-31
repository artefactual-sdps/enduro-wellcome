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
/**
 * PreservationTask describes a preservation action task. (default view)
 * @export
 * @interface EnduroPackagePreservationTaskResponseBody
 */
export interface EnduroPackagePreservationTaskResponseBody {
    /**
     * 
     * @type {Date}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    completedAt?: Date;
    /**
     * 
     * @type {number}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    note?: string;
    /**
     * 
     * @type {number}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    preservationActionId?: number;
    /**
     * 
     * @type {Date}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    startedAt: Date;
    /**
     * 
     * @type {string}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    status: EnduroPackagePreservationTaskResponseBodyStatusEnum;
    /**
     * 
     * @type {string}
     * @memberof EnduroPackagePreservationTaskResponseBody
     */
    taskId: string;
}


/**
 * @export
 */
export const EnduroPackagePreservationTaskResponseBodyStatusEnum = {
    Unspecified: 'unspecified',
    InProgress: 'in progress',
    Done: 'done',
    Error: 'error',
    Queued: 'queued',
    Pending: 'pending'
} as const;
export type EnduroPackagePreservationTaskResponseBodyStatusEnum = typeof EnduroPackagePreservationTaskResponseBodyStatusEnum[keyof typeof EnduroPackagePreservationTaskResponseBodyStatusEnum];


/**
 * Check if a given object implements the EnduroPackagePreservationTaskResponseBody interface.
 */
export function instanceOfEnduroPackagePreservationTaskResponseBody(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "startedAt" in value;
    isInstance = isInstance && "status" in value;
    isInstance = isInstance && "taskId" in value;

    return isInstance;
}

export function EnduroPackagePreservationTaskResponseBodyFromJSON(json: any): EnduroPackagePreservationTaskResponseBody {
    return EnduroPackagePreservationTaskResponseBodyFromJSONTyped(json, false);
}

export function EnduroPackagePreservationTaskResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroPackagePreservationTaskResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'completedAt': !exists(json, 'completed_at') ? undefined : (new Date(json['completed_at'])),
        'id': json['id'],
        'name': json['name'],
        'note': !exists(json, 'note') ? undefined : json['note'],
        'preservationActionId': !exists(json, 'preservation_action_id') ? undefined : json['preservation_action_id'],
        'startedAt': (new Date(json['started_at'])),
        'status': json['status'],
        'taskId': json['task_id'],
    };
}

export function EnduroPackagePreservationTaskResponseBodyToJSON(value?: EnduroPackagePreservationTaskResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'completed_at': value.completedAt === undefined ? undefined : (value.completedAt.toISOString()),
        'id': value.id,
        'name': value.name,
        'note': value.note,
        'preservation_action_id': value.preservationActionId,
        'started_at': (value.startedAt.toISOString()),
        'status': value.status,
        'task_id': value.taskId,
    };
}

