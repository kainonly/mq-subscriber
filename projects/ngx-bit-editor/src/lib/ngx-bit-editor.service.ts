import {Injectable} from '@angular/core';

@Injectable()
export class NgxBitEditorService {
  exec(id: string, value: string): boolean {
    return document.execCommand(id, false, value);
  }
}
