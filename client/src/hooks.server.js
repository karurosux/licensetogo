import { newPocketBase } from '$lib/utils/pb';

/**
 * Handles the server-side hooks
 * @param {import('@sveltejs/kit').Handle} handle
 */
// @ts-ignore
export const handle = async ({ event, resolve }) => {
	const pb = newPocketBase();

	pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

	try {
		if (pb.authStore.isValid) {
			await pb.collection('users').authRefresh();
		}
	} catch (_) {
		// clear the auth store on failed refresh
		pb.authStore.clear();
	}

	event.locals.pb = pb;
	event.locals.user = pb.authStore.record;

	const response = await resolve(event);

	// send back the default 'pb_auth' cookie to the client with the latest store state
	response.headers.set('set-cookie', pb.authStore.exportToCookie({ httpOnly: false }));

	return response;
};
