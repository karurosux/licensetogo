import { catchPromise } from '$lib/utils/catch-promise';
import { redirect } from '@sveltejs/kit';

export const actions = {
	login: async ({ request, locals }) => {
		/**
		 * @type {import("$lib/utils/pb").pb}
		 */
		const pb = locals.pb;
		const body = await request.formData();
		const email = body.get('email');
		const password = body.get('password');

		if (!email || !password) {
			return {
				failed: true
			};
		}

		const res = await catchPromise(pb.collection('users').authWithPassword(email, password));

		if (res.ok) {
			redirect(303, '/license-manager');
		}

		return {
			failed: res.ok
		};
	}
};
