import type {ApiResponse} from "@/utils/request";
import service from "@/utils/request";

export interface Murmur {
    id: number;
    content: string;
    created_at: string;
    updated_at: string;
}

export const murmurAll = (): Promise<ApiResponse<Murmur[]>> => {
    return service({
        url: '/murmur/all',
        method: 'get',
    })
}

export interface MurmurCreateRequest {
    content: string;
}

export const murmurCreate = (data: MurmurCreateRequest): Promise<ApiResponse<undefined>> => {
    return service({
        url: '/murmur/create',
        method: 'post',
        data: data,
    })
}

export interface MurmurDeleteRequest {
    ids: number[];
}

export const murmurDelete = (data: MurmurDeleteRequest): Promise<ApiResponse<undefined>> => {
    return service({
        url: '/murmur/delete',
        method: 'delete',
        data: data,
    })
}
