import { LocationUpdate } from '@smartems/runtime';
import { actionCreatorFactory } from 'app/core/redux';

export const updateLocation = actionCreatorFactory<LocationUpdate>('UPDATE_LOCATION').create();
