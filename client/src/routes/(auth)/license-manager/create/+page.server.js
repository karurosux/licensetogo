import { setLicenceTransference } from '$lib/context/license-transfer.js';
import { catchPromise } from '$lib/utils/catch-promise';
import { redirect } from '@sveltejs/kit';

export const actions = {
	create: async ({ request, locals }) => {
		/**
		 * @type {import("$lib/utils/pb").pb}
		 */
		const pb = locals.pb;
		const body = await request.formData();
		const name = body.get('name');
		const expires = body.get('expires');

		if (!name) {
			return { failed: true };
		}

		const requestBody = {
			name: name
		};

		if (expires) {
			requestBody.expires = expires;
		}

		const res = await catchPromise(
			pb.send('/api/license', {
				method: 'POST',
				body: JSON.stringify(requestBody)
			})
		);

		if (!res.ok) {
			return { failed: true };
		}

		const code = setLicenceTransference(res.data?.key);

		redirect(303, '/license-manager/create/success?m=' + code);

		return { failed: false };
	}
};
