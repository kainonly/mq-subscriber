import {NgModule} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {NgxBitEditorComponent} from './ngx-bit-editor.component';
import {NgxBitEditorService} from './ngx-bit-editor.service';

@NgModule({
  imports: [
    FormsModule
  ],
  declarations: [NgxBitEditorComponent],
  exports: [NgxBitEditorComponent],
  providers: [NgxBitEditorService]
})
export class NgxBitEditorModule {
}
