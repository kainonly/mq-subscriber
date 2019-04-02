import {NgModule} from '@angular/core';
import {NgxLiteEditorComponent} from './ngx-lite-editor.component';
import {CommonModule} from '@angular/common';
import {NgZorroAntdModule} from 'ng-zorro-antd';

@NgModule({
  imports: [
    CommonModule,
    NgZorroAntdModule
  ],
  declarations: [
    NgxLiteEditorComponent
  ],
  exports: [
    NgxLiteEditorComponent
  ]
})
export class NgxLiteEditorModule {
}
