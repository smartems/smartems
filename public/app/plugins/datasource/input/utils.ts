import { toDataFrame, DataFrameDTO, toCSV } from '@smartems/data';

export function dataFrameToCSV(dto?: DataFrameDTO[]) {
  if (!dto || !dto.length) {
    return '';
  }
  return toCSV(dto.map(v => toDataFrame(v)));
}
