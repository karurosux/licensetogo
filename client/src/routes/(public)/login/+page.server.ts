import { Collections } from '$lib/models/generated/pb-models.js';
import { catchPromise } from '$lib/utils/catch-promise';

export const actions = {
	login: async ({ request, locals }) => {
		const body = await request.formData();
		const email = body.get('email');
		const password = body.get('password');

		if (!email || !password) {
			return {
				failed: true
			};
		}

		const res = await catchPromise(
			locals.pb.collection(Collections.Users).authWithPassword(email as string, password as string)
		);

		return {
			failed: !res.ok
		};
	}
};
