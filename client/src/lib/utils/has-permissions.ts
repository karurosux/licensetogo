import type {
	RolesPermissionOptions,
	RolesResponse,
	UsersResponse
} from '$lib/models/generated/pb-models';

export const hasPermissions = (
	user: UsersResponse<{ role: RolesResponse }>,
	...permisisons: RolesPermissionOptions[]
) => {
	return permisisons.some((p) => user.expand?.role.permission.includes(p));
};
