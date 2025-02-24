import { goto } from '$app/navigation';

/**
 * This is a helper function to replace the current URL state with the given
 * query parameters. This is useful when you want to preserve the current
 * page state but change the query parameters.
 *
 * @param {Record<string, string | number>} values - The query parameters to replace with.
 */
export const replaceStateWithQuery = (values) => {
	const searchParams = new URLSearchParams();
	for (let [k, v] of Object.entries(values)) {
		if (!!v) {
			searchParams.set(k, v.toString());
		} else {
			searchParams.delete(k);
		}
	}
	goto(Object.keys(values).length ? `?${searchParams.toString()}` : '');
};
