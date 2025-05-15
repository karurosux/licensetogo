import { writable } from 'svelte/store';

export type BreadcrumbItem = {
	href?: string;
	label: string;
};

export const breadcrumbs = writable<BreadcrumbItem[]>([]);
