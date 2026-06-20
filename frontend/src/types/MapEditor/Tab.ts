export enum SubTabAction {
    Upload      = 'upload',
    NodeCreate  = 'nodeCreate',
    NodeDelete  = 'nodeDelete',
    RouteCreate = 'routeCreate',
    RouteDelete = 'routeDelete',
}

export interface SubTab {
    name: string;
    icon?: string | null;
    action?: SubTabAction;
}

export interface Tab {
    name: string;
    icon?: string | null;
    subTabs?: SubTab[];
}