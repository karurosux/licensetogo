import { Collections, type LicenseResponse } from '$lib/models/generated/pb-models.js';
import { catchPromise } from '$lib/utils/catch-promise.js';

export const load = async ({ locals, url }) => {
	const page = url.searchParams.get('page') || 1;
	const perPage = url.searchParams.get('perPage') || 5;
	const filter = url.searchParams.get('filter');

	const res = await catchPromise(
		locals.pb.collection(Collections.License).getList<LicenseResponse>(+page, +perPage, {
			filter: filter ? `name~"%${filter.toLowerCase()}%"` : undefined
		})
	);

	return {
		licenses: res.data,
		ok: res.ok,
		error: res.error && {
			message: res.error.message
		}
	};
};
