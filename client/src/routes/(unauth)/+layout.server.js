import { redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	// @ts-ignore
	if (locals.user) {
		redirect(303, '/license-manager');
	}
};
