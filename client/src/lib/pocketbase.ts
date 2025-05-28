import PocketBase from 'pocketbase';
import { env } from '$env/dynamic/public';
import type { TypedPocketBase } from './models/generated/pb-models';

export const createPocketbase = () => {
	return new PocketBase(env?.PUBLIC_API_URL || 'http://localhost:8090') as TypedPocketBase;
};

export const pb = createPocketbase();
