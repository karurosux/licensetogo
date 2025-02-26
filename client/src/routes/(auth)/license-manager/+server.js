import { json, text } from '@sveltejs/kit';

export async function POST({ request, locals }) {
	/**
	 * @type {import('$lib/utils/pb.js').pb}
	 */
	const pb = locals.pb;
	const body = await request.json();

	switch (body.action) {
		case 'set-active':
			await pb.collection('license').update(body.id, {
				active: body.value
			});
			return text('');
	}
}
