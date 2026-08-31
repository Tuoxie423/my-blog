import type {ApiResponse} from "@/utils/request";
import service from "@/utils/request";

export interface HotItem {
    index: number;
    title: string;
    url: string;
    hot_value: string;
}

export interface HotPlatform {
    type: string;
    name: string;
    icon: string;
    list: HotItem[];
}

export const hotAll = (): Promise<ApiResponse<HotPlatform[]>> => {
    return service({
        url: '/hot/all',
        method: 'get',
    })
}
