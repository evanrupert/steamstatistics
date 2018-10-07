import { BrowserModule } from '@angular/platform-browser'
import { NgModule } from '@angular/core'
import { StoreModule } from '@ngrx/store'
import { dataReducer } from './reducers/data.reducer'

import { AppComponent } from './app.component'
import { VanityUrlInputComponent } from './vanity-url-input/vanity-url-input.component'
import { PlaytimeGraphComponent } from './playtime-graph/playtime-graph.component'

@NgModule({
  declarations: [
    AppComponent,
    VanityUrlInputComponent,
    PlaytimeGraphComponent
  ],
  imports: [
    BrowserModule,
    StoreModule.forRoot({ data: dataReducer })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
