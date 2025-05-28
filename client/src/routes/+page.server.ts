import { redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	redirect(307, locals?.user ? '/home' : '/login');
};
