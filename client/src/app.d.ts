import type {
	RolesResponse,
	TypedPocketBase,
	UsersResponse
} from '$lib/models/generated/pb-models';
import PocketBase from 'pocketbase';

// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			pb: TypedPocketBase;
			user: UsersResponse<{ role: RolesResponse }>;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
