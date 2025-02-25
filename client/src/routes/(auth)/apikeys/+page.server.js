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
		fields: 'id,name,active,created',
		order: '-created'
	};

	if (filter) {
		options.filter = `name ~ "%${filter}%"`;
	}

	const res = pb.collection('apikey').getList(offset + 1, limit, options);

	return {
		license: res,
		query: {
			filter,
			offset,
			limit
		}
	};
};
