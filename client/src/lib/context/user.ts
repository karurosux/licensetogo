import type { RolesResponse, UsersResponse } from '$lib/models/generated/pb-models';
import { getContext, setContext } from 'svelte';
import { writable, type Writable } from 'svelte/store';

const userKey = Symbol('user');

export function setUserContext(user: Writable<UsersResponse<{ role: RolesResponse }>>) {
	setContext(userKey, user);
}

export function getUserContext(): Writable<UsersResponse<{ role: RolesResponse }>> {
	return getContext(userKey);
}

export let currentUser = writable<UsersResponse<{ role: RolesResponse }> | null>(null);
