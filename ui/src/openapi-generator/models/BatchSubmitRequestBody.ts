/* tslint:disable */
/* eslint-disable */
/**
 * Enduro API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface BatchSubmitRequestBody
 */
export interface BatchSubmitRequestBody {
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    path: string;
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    pipeline: string;
    /**
     * 
     * @type {string}
     * @memberof BatchSubmitRequestBody
     */
    processingConfig?: string;
}

export function BatchSubmitRequestBodyFromJSON(json: any): BatchSubmitRequestBody {
    return BatchSubmitRequestBodyFromJSONTyped(json, false);
}

export function BatchSubmitRequestBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): BatchSubmitRequestBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'path': json['path'],
        'pipeline': json['pipeline'],
        'processingConfig': !exists(json, 'processing_config') ? undefined : json['processing_config'],
    };
}

export function BatchSubmitRequestBodyToJSON(value?: BatchSubmitRequestBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'path': value.path,
        'pipeline': value.pipeline,
        'processing_config': value.processingConfig,
    };
}


