/**
 * Handles the server-side layout
 * @param {import('@sveltejs/kit').Load} event - The server-side layout load event
 */

// @ts-ignore
export const load = async ({ locals }) => {
	return {
		user: locals.user
	};
};
