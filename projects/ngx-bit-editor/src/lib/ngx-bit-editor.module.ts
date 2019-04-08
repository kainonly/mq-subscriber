import {NgModule} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {NgxBitEditorComponent} from './ngx-bit-editor.component';

@NgModule({
  imports: [
    FormsModule
  ],
  declarations: [NgxBitEditorComponent],
  exports: [NgxBitEditorComponent],
})
export class NgxBitEditorModule {
}
