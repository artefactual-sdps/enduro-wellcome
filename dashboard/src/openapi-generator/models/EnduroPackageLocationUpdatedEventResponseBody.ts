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
 * EnduroPackage-Location-Updated-EventResponseBody result type (default view)
 * @export
 * @interface EnduroPackageLocationUpdatedEventResponseBody
 */
export interface EnduroPackageLocationUpdatedEventResponseBody {
    /**
     * Identifier of package
     * @type {number}
     * @memberof EnduroPackageLocationUpdatedEventResponseBody
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof EnduroPackageLocationUpdatedEventResponseBody
     */
    locationId: string;
}

/**
 * Check if a given object implements the EnduroPackageLocationUpdatedEventResponseBody interface.
 */
export function instanceOfEnduroPackageLocationUpdatedEventResponseBody(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "locationId" in value;

    return isInstance;
}

export function EnduroPackageLocationUpdatedEventResponseBodyFromJSON(json: any): EnduroPackageLocationUpdatedEventResponseBody {
    return EnduroPackageLocationUpdatedEventResponseBodyFromJSONTyped(json, false);
}

export function EnduroPackageLocationUpdatedEventResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EnduroPackageLocationUpdatedEventResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'locationId': json['location_id'],
    };
}

export function EnduroPackageLocationUpdatedEventResponseBodyToJSON(value?: EnduroPackageLocationUpdatedEventResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'location_id': value.locationId,
    };
}

