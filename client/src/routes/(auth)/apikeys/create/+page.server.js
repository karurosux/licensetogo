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

		if (!name) {
			return { failed: true };
		}

		const res = await catchPromise(
			pb.send('/api/apikey', {
				method: 'POST',
				body: JSON.stringify({
					name: name
				})
			})
		);

		if (!res.ok) {
			return { failed: true };
		}

		const code = setLicenceTransference(res.data?.key);

		return redirect(303, '/apikeys/create/success?m=' + code);
	}
};
