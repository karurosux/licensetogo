import { getContext, setContext } from 'svelte';

const userKey = Symbol('user');

/**
 * Set the user context
 * @param {import('svelte/store').Writable<import('pocketbase').AuthRecord>} user
 */
export function setUserContext(user) {
	setContext(userKey, user);
}

/**
 * Get the user context
 * @returns {import('svelte/store').Writable<import('pocketbase').AuthRecord>}
 */
export function getUserContext() {
	return getContext(userKey);
}
