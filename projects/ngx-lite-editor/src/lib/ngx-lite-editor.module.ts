import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule} from '@angular/forms';
import {NgZorroAntdModule} from 'ng-zorro-antd';
import {NgxLiteEditorComponent} from './ngx-lite-editor.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    NgZorroAntdModule,
  ],
  declarations: [
    NgxLiteEditorComponent,
  ],
  exports: [
    NgxLiteEditorComponent
  ]
})
export class NgxLiteEditorModule {
}
