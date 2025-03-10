import { catchPromise } from '$lib/utils/catch-promise';
import { getPBFromLocals } from '$lib/utils/pb';

export const actions = {
	update: async ({ request, locals }) => {
		const pb = getPBFromLocals(locals);
		const body = await request.formData();
		const email = body.get('email');
		const password = body.get('password');
		const oldPassword = body.get('oldPassword');

		const res = await catchPromise(
			pb.send('api/settings/credentials', {
				method: 'PUT',
				body: {
					email: email,
					password: password,
					oldPassword: oldPassword
				}
			})
		);
		console.log('password updated -> ', res);
		return {
			success: res.ok,
			error: !!res.error
		};
	}
};
