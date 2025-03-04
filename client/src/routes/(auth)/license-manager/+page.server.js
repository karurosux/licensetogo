import { getPBFromLocals } from '$lib/utils/pb';

/**
 * @type {import("./$types").PageServerLoad}
 */
export const load = async ({ url, locals }) => {
	const query = new URL(url).searchParams;
	const limit = +(query.get('limit') || 10);
	const offset = +(query.get('offset') || 0);
	const filter = query.get('filter') || '';
	/**
	 * @type {import("$lib/utils/pb").pb}
	 */
	const pb = locals.pb;
	const options = {
		fields: 'id,name,permissions,metadata,active,expires,lastused',
		order: '-created'
	};

	if (filter) {
		options.filter = `name ~ "%${filter}%"`;
	}

	const res = pb.collection('license').getList(offset + 1, limit, options);

	return {
		license: res,
		query: {
			filter,
			offset,
			limit
		}
	};
};

export const actions = {
	setActive: async ({ request, locals }) => {
		const body = await request.formData();
		const pb = getPBFromLocals(locals);
		await pb.collection('license').update(body.get('id'), { active: body.get('value') });
		return {
			success: true
		};
	},
	delete: async ({ request, locals }) => {
		const body = await request.formData();
		const pb = getPBFromLocals(locals);
		await pb.collection('license').delete(body.get('id'));

		return {
			success: true
		};
	}
};
