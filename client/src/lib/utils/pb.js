import PocketBase from 'pocketbase';
import { env } from '$env/dynamic/public';

export function newPocketBase() {
	return new PocketBase(env?.PUBLIC_API_URL || 'http://localhost:8090');
}

export const pb = newPocketBase();

/**
 * @param {any} locals
 * @return {pb}
 */
export const getPBFromLocals = (locals) => {
	return locals.pb;
};
