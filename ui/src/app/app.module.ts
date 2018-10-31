import { BrowserModule } from '@angular/platform-browser'
import { NgModule } from '@angular/core'
import { StoreModule } from '@ngrx/store'
import { dataReducer } from './reducers/data.reducer'

import { AppComponent } from './app.component'
import { VanityUrlInputComponent } from './vanity-url-input/vanity-url-input.component'
import { PlaytimeGraphComponent } from './playtime-graph/playtime-graph.component'
import { ChartsModule } from 'ng2-charts'
import { MatProgressBarModule } from '@angular/material'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'

@NgModule({
  declarations: [
    AppComponent,
    VanityUrlInputComponent,
    PlaytimeGraphComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    StoreModule.forRoot({ data: dataReducer }),
    ChartsModule,
    MatProgressBarModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
