import { Collections, type LicenseResponse } from '$lib/models/generated/pb-models.js'
import { catchPromise } from '$lib/utils/catch-promise.js'

export const load = async ({ locals, url }) => {
	const page = url.searchParams.get('page') || 1
	const perPage = url.searchParams.get('perPage') || 5
	const filter = url.searchParams.get('filter')

	const res = await catchPromise(
		locals.pb.collection(Collections.License).getList<LicenseResponse>(+page, +perPage, {
			filter: filter ? `name~"%${filter.toLowerCase()}%"` : undefined,
		}),
	)

	return {
		licenses: res.data,
		ok: res.ok,
		error: res.error && {
			message: res.error.message,
		},
	}
}

export const actions = {
	create: async ({ request, locals }) => {
		const body = await request.formData()
		const name = body.get('name')
		const expires = body.get('expires')

		if (!name) {
			return {
				failed: true,
				error: 'missing_required_fields',
			}
		}

		const bodyObject = {
			name: name as string,
			expires: expires || null,
			lastused: null,
			metadata: null,
			permissions: null,
		}

		const res = await catchPromise(
			locals.pb.send('/api/license', {
				method: 'POST',
				body: bodyObject,
			}),
		)

		return {
			failed: !res.ok,
			error: res.error?.message,
		}
	},
	delete: async ({ request, locals }) => {
		const body = await request.formData()
		const id = body.get('id')

		if (!id) {
			return {
				failed: true,
				error: 'missing_required_fields',
			}
		}

		const res = await catchPromise(locals.pb.collection(Collections.License).delete(id as string))

		return {
			failed: !res.ok,
			error: res.error?.message,
		}
	},
	active: async ({ request, locals }) => {
		const body = await request.formData()
		const id = body.get('id')
		const value = body.get('value')

		if (!id) {
			return {
				failed: true,
				error: 'missing_required_fields',
			}
		}

		const res = await catchPromise(
			locals.pb.collection(Collections.License).update(id as string, {
				active: value,
			}),
		)

		return {
			failed: !res.ok,
			error: res?.error?.message,
		}
	},
}
