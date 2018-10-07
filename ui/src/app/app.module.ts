import { BrowserModule } from '@angular/platform-browser'
import { NgModule } from '@angular/core'
import { StoreModule } from '@ngrx/store'
import { dataReducer } from './reducers/data.reducer'

import { AppComponent } from './app.component'
import { VanityUrlInputComponent } from './vanity-url-input/vanity-url-input.component'

@NgModule({
  declarations: [
    AppComponent,
    VanityUrlInputComponent
  ],
  imports: [
    BrowserModule,
    StoreModule.forRoot({ data: dataReducer })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
