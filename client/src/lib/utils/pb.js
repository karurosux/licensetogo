import PocketBase from 'pocketbase';

export function newPocketBase() {
	return new PocketBase('http://localhost:8090');
}

export const pb = newPocketBase();
