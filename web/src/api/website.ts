import type {ApiResponse} from "@/utils/request";
import service from "@/utils/request";
import type {Website} from "@/api/config";

export const websiteInfo = (): Promise<ApiResponse<Website>> => {
    return service({
        url: '/website/info',
        method: 'get',
    })
}

export interface FooterLink {
    title: string;
    link: string;
}

export const websiteFooterLink = (): Promise<ApiResponse<FooterLink[]>> => {
    return service({
        url: '/website/footerLink',
        method: 'get',
    })
}

export const websiteCreateFooterLink = (data: FooterLink): Promise<ApiResponse<undefined>> => {
    return service({
        url: '/website/createFooterLink',
        method: 'post',
        data: data,
    })
}

export const websiteDeleteFooterLink = (data: FooterLink): Promise<ApiResponse<undefined>> => {
    return service({
        url: '/website/deleteFooterLink',
        method: 'delete',
        data: data,
    })
}

export const websiteYiyan = (): Promise<ApiResponse<string>> => {
    return service({
        url: '/website/yiyan',
        method: 'get',
    })
}
