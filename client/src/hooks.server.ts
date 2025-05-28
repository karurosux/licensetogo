import { Collections } from '$lib/models/generated/pb-models';
import { createPocketbase } from '$lib/pocketbase';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const pb = createPocketbase();
	const cookieString = event.request.headers.get('cookie') || '';

	pb.authStore.loadFromCookie(cookieString);

	let refreshRes: { token: string; record: any } | null = null;

	try {
		if (pb.authStore.isValid) {
			refreshRes = await pb.collection(Collections.Users).authRefresh({
				expand: 'role'
			});
		}
	} catch (_) {
		// clear the auth store on failed refresh
		pb.authStore.clear();
	}

	event.locals.pb = pb;
	event.locals.user = refreshRes?.record;

	const response = await resolve(event, {
		transformPageChunk: ({ html }) => {
			return html.replace(
				'<html',
				`<html class="${cookieString.includes('theme=light') ? '' : 'dark'}"`
			);
		}
	});

	// send back the default 'pb_auth' cookie to the client with the latest store state
	response.headers.set('set-cookie', pb.authStore.exportToCookie({ httpOnly: false }));

	return response;
};
