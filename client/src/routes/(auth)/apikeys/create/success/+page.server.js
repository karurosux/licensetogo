import { getLicenseTransference } from '$lib/context/license-transfer.js';
import { error } from '@sveltejs/kit';

export const load = ({ url }) => {
	const query = new URL(url).searchParams;
	const code = query.get('m');

	if (!code) {
		throw error(400, 'Missing license code');
	}

	const license = getLicenseTransference(code);

	if (!license) {
		throw error(404, 'License not found');
	}

	return {
		license: license
	};
};
