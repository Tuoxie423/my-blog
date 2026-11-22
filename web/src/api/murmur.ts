import type {ApiResponse} from "@/utils/request";
import service from "@/utils/request";

export interface Murmur {
    id: number;
    content: string;
    created_at: string;
    updated_at: string;
    likes: number;
}

export interface MurmurWithLiked extends Murmur {
    is_liked: boolean;
}

export interface MurmurLikeResult {
    is_liked: boolean;
    likes: number;
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

export interface MurmurLikeRequest {
    murmur_id: number;
}

export const murmurLike = (data: MurmurLikeRequest): Promise<ApiResponse<MurmurLikeResult>> => {
    return service({
        url: '/murmur/like',
        method: 'post',
        data: data,
    })
}

export const murmurLiked = (): Promise<ApiResponse<number[]>> => {
    return service({
        url: '/murmur/liked',
        method: 'get',
    })
}
